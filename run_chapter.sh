#!/bin/bash

# Go语法学习 - 运行指定章节

if [ $# -eq 0 ]; then
    echo "使用方法: $0 <章节号>"
    echo ""
    echo "可用章节："
    echo "  1 - 基础语法"
    echo "  2 - 数据类型"
    echo "  3 - 控制流"
    echo "  4 - 函数"
    echo "  5 - 结构体与方法"
    echo "  6 - 接口"
    echo "  7 - 并发编程"
    echo "  8 - 包管理"
    echo "  9 - 错误处理"
    echo "  10 - 实战练习"
    echo ""
    echo "示例: $0 1"
    exit 1
fi

chapter_num=$1

case $chapter_num in
    1)
        file="01_basics/hello.go"
        title="基础语法"
        ;;
    2)
        file="02_types/types.go"
        title="数据类型"
        ;;
    3)
        file="03_control_flow/control.go"
        title="控制流"
        ;;
    4)
        file="04_functions/functions.go"
        title="函数"
        ;;
    5)
        file="05_structs_methods/structs.go"
        title="结构体与方法"
        ;;
    6)
        file="06_interfaces/interfaces.go"
        title="接口"
        ;;
    7)
        file="07_concurrency/concurrency.go"
        title="并发编程"
        ;;
    8)
        file="08_packages/main.go"
        title="包管理"
        ;;
    9)
        file="09_error_handling/errors.go"
        title="错误处理"
        ;;
    10)
        file="10_practice/practice.go"
        title="实战练习"
        ;;
    *)
        echo "❌ 无效的章节号: $chapter_num"
        echo "   请输入 1-10 之间的数字"
        exit 1
        ;;
esac

echo "📖 第${chapter_num}章：$title"
echo "========================================"
echo "运行文件：$file"
echo ""

if [ -f "$file" ]; then
    go run "$file"
    if [ $? -eq 0 ]; then
        echo ""
        echo "✅ 第${chapter_num}章运行成功"
    else
        echo ""
        echo "❌ 第${chapter_num}章运行失败"
    fi
else
    echo "❌ 文件不存在：$file"
fi
