#!/bin/bash

# 检查是否提供了题号
if [ -z "$1" ]; then
  echo "Usage: $0 <slug>"
  echo "Example: $0 two-sum"
  exit 1
fi

slug="$1"

# 发送 GraphQL 请求获取题目信息
response=$(curl -s 'https://leetcode.com/graphql' \
  -H 'Content-Type: application/json' \
  --data-binary "{\"query\":\"query { question(titleSlug: \\\"$slug\\\") { questionFrontendId title content codeSnippets { lang langSlug code } } }\"}")

# 解析标题
questionId=$(echo "$response" | jq -r '.data.question.questionFrontendId')
title=$(echo "$response" | jq -r '.data.question.title')
# 格式化描述信息
desc=$(echo "$response" | jq -r '.data.question.content' |
 sed 's/<sup>/^/g' | # 处理指数符号
 sed 's/<[^>]*>//g' | # 去掉HTML标签
 sed -e 's/&lt;/</g' -e 's/&gt;/>/g' -e 's/&amp;/\&/g' -e 's/&quot;/"/g' -e "s/&apos;/'/g" -e 's/&nbsp;/ /g' | # 处理HTML实体转换
 sed -e 's/\t//g' |
 grep -vE '^\s*$' | # 删除空行
 sed -e 's/^\(Example [0-9]:\)$/\n\1/g' | # 在每个Example之前添加空行
 sed -e 's/^\(Constraints:\)$/\n\1/g' | # 在Constraints之前添加空行
 sed -e 's/^\(Follow up:\)$/\n\1/g' | # 在Follow up:之前添加空行
 sed 's/^\(.*\)$/\/\/ \1/g') # 添加注释符
# 解析默认代码模板（以 Python 为例）
code=$(echo "$response" | jq -r '.data.question.codeSnippets[] | select(.langSlug=="golang") | .code')

# 判断文件是否存在
filename=$(echo "$questionId""_$slug.go" | sed 's/-/_/g')

if [ -e "$filename" ]; then
  echo "文件已存在，请勿重复创建"
  exit 1
fi

# 将内容写入到文件中
common="package leetcode

// https://leetcode.com/problems/$slug
//"
echo "$common" > $filename
echo "$desc" >> $filename
echo "$code" >> $filename
