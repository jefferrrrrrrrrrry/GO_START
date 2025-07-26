#!/bin/bash

# Go语法学习 - 运行所有示例

echo "🚀 Go语法快速学习 - 开始运行所有示例"
echo "========================================"

# 检查是否安装了Go
if ! command -v go &> /dev/null; then
    echo "❌ 错误：未找到Go安装。请先安装Go语言。"
    echo "   下载地址：https://golang.org/dl/"
    exit 1
fi

echo "✅ Go版本：$(go version)"
echo ""

chapters=(
    "01_basics/hello.go:基础语法"
    "02_types/types.go:数据类型"
    "03_control_flow/control.go:控制流"
    "04_functions/functions.go:函数"
    "05_structs_methods/structs.go:结构体与方法"
    "06_interfaces/interfaces.go:接口"
    "07_concurrency/concurrency.go:并发编程"
    "08_packages/main.go:包管理"
    "09_error_handling/errors.go:错误处理"
    "10_practice/practice.go:实战练习"
)

for chapter in "${chapters[@]}"; do
    IFS=':' read -r file title <<< "$chapter"
    
    echo "📖 第${file%%/*}章：$title"
    echo "----------------------------------------"
    echo "运行文件：$file"
    echo ""
    
    if [ -f "$file" ]; then
        # 运行Go文件
        if [[ "$file" == "10_practice/practice.go" ]]; then
            echo "⚠️  注意：实战练习是交互式程序，请手动运行："
            echo "   go run $file"
        else
            go run "$file"
            if [ $? -eq 0 ]; then
                echo "✅ 运行成功"
            else
                echo "❌ 运行失败"
            fi
        fi
    else
        echo "❌ 文件不存在：$file"
    fi
    
    echo ""
    echo "========================================"
    echo ""
    
    # 除了最后一个练习，其他都暂停一下
    if [[ "$file" != "10_practice/practice.go" ]]; then
        read -p "按Enter键继续下一章..." -r
        echo ""
    fi
done

echo "🎉 所有示例运行完成！"
echo ""
echo "💡 学习建议："
echo "   1. 重新运行感兴趣的章节: go run 章节路径/文件名.go"
echo "   2. 尝试修改代码，观察输出变化"
echo "   3. 完成实战练习：go run 10_practice/practice.go"
echo "   4. 阅读详细学习指南：查看 LEARNING_GUIDE.md"
echo ""
echo "🚀 开始你的Go语言编程之旅吧！"
