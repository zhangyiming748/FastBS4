#!/bin/bash

# 下载当前目录中每一个txt文件中每一个URL的脚本
# 文件名从URL最后一个路径分隔符后的部分提取，并进行URL解码

urldecode() {
    local encoded="$1"
    printf '%b' "${encoded//%/\\x}"
}

for txtfile in *.txt; do
    [ -f "$txtfile" ] || continue
    echo "处理文件: $txtfile"
    while IFS= read -r url || [ -n "$url" ]; do
        # 跳过空行和注释
        [ -z "$url" ] && continue
        [[ "$url" =~ ^# ]] && continue

        # 提取文件名（最后一个/之后的部分）
        filename="${url##*/}"
        # URL解码
        filename="$(urldecode "$filename")"

        echo "下载: $url -> $filename"
        wget -O "$filename" "$url"
    done < "$txtfile"
done
