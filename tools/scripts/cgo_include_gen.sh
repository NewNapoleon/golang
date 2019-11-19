#!/bin/bash
#
# Perform auto import of non go files.
# [go mod vendor is not copying all files to the vendor folder](https://github.com/golang/go/issues/27832)

set -o pipefail
set -o errexit
set -o nounset
# set -o xtrace

# 获取输入参数
THIS_BASE_PARAM="$*"
# 获取当前脚本的相对路径文件名称
THIS_BASH_FILE="${BASH_SOURCE-$0}"
# 获取当前脚本的相对路径目录
THIS_BASH_FILE_REF_DIR=$(dirname "${THIS_BASH_FILE}")
# 获取当前脚本的绝对路径目录
THIS_BASH_FILE_ABS_DIR=$(
  cd "${THIS_BASH_FILE_REF_DIR}" || exit
  pwd
)
# 获取当前脚本的名称
THIS_BASH_FILE_BASE_NAME=$(basename "${THIS_BASH_FILE}")
# 获取当前脚本绝对路径
THIS_BASH_FILE_ABS_PATH="${THIS_BASH_FILE_ABS_DIR}/${THIS_BASH_FILE_BASE_NAME}"
# 备份当前路径
STACK_ABS_DIR=$(pwd)
# 路径隔离
#cd "${THIS_BASH_FILE_ABS_DIR}" 1>/dev/null 2>&1 || exit

# This script rebuilds the generated code for the protocol buffers.
# To run this you will need protoc and goprotobuf installed;
# see https://github.com/golang/protobuf for instructions.
# You also need Go and Git installed.
# masking all out put info!

log_date() {
  echo "[$(date +'%Y-%m-%dT%H:%M:%S%z')]: $*" >&1
}

err_date() {
  echo "[$(date +'%Y-%m-%dT%H:%M:%S%z')]: $*" >&2
}

function set_default_var_param() {
  g_cfg_force_mode=""         # #覆盖前永不提示-f

  g_pkg_root_dirs=()          # 扫描并生成包的根目录列表
  g_pkg_root_import_prefix="" # 包根目录的前缀，如github.com/searKing/golang
}

# include
pushd "${THIS_BASH_FILE_ABS_DIR}" 1>/dev/null 2>&1 || exit
. "./lib/check_tools.sh"
popd 1>/dev/null 2>&1 || exit

set_default_var_param #设置默认变量参数
while getopts "fp:h" opt; do
  case $opt in
  f)
    #覆盖前永不提示
    g_cfg_force_mode=1
    ;;
  h)
    usage
    exit 1
    ;;
  p)
    g_pkg_root_import_prefix="${OPTARG}"
    ;;
  ?)
    log_error "${LINENO}:$opt is Invalid"
    ;;
  *) ;;

  esac
done
#去除options参数
shift $((OPTIND - 1))

if [ "$#" -lt 1 ]; then
  cat <<HELPEOF
use option -h to get more log_information.
HELPEOF
  exit 1
fi
g_pkg_root_dirs=("$@")

readonly g_cfg_force_mode
readonly g_pkg_root_dirs
readonly g_pkg_root_import_prefix

#去除options参数
#shift n表示把第n+1个参数移到第1个参数, 即命令结束后$1的值等于$n+1的值
#shift ${#g_pkg_root_dirs[@]}

# Sanity check that the right tools are accessible.
#check_tools()
echo "dirs to be scanned : ${g_pkg_root_dirs}"
for pkg_root_dir in ${g_pkg_root_dirs}; do
  printf "%s scanning" "${pkg_root_dir}"

  pkg_root_base_name="$(basename "${pkg_root_dir}")"
  readonly pkg_root_base_name
  pkg_root_go="${pkg_root_base_name}.cgo.go"
  readonly pkg_root_go

  pushd "${pkg_root_dir}" 1>/dev/null 2>&1 || exit

  cat <<END >"${pkg_root_go}"
// Code generated by "$0"; DO NOT EDIT.

package ${pkg_root_base_name}

END

  echo ""
  find "." -type d -print0 | while read -r -d $'\0' sub_dir; do
    if [[ "${sub_dir}" = "." ]]; then
        continue
    fi
    pkg_base_name="$(basename "${sub_dir}")"
    pkg_dir="$(dirname "${sub_dir}")"
    pkg_import_path="${sub_dir#./}"
    if [[ -n "${g_pkg_root_import_prefix}" ]]; then
      pkg_import_path="${g_pkg_root_import_prefix%/}/${pkg_import_path}"
    fi

    pushd "${sub_dir}" 1>/dev/null 2>&1 || exit

    pkg_go="${pkg_base_name}.cgo.go"
    printf "\r\033[K%s generating" "${sub_dir}"
    cat <<END >"${pkg_go}"
// Code generated by "$0"; DO NOT EDIT.

package ${pkg_base_name}

END
    printf "\r\033[K%s generated" "${sub_dir}"
    popd 1>/dev/null 2>&1 || exit

    cat <<END >>"${pkg_root_go}"
import _ "${pkg_import_path}"
END
  done
  printf "\r\033[K"
  printf "\033[1A"
  printf "\r\033[K%s scanned\n" "${pkg_root_dir}"

  popd 1>/dev/null 2>&1 || exit
done

printf "\r\033[Kcgo-include-gen done...\n"