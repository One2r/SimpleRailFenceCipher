package main

import (
	"os"
	"strconv"
	"fmt"

	"github.com/codegangsta/cli"
)


func main() {
	app := cli.NewApp()
	app.Name = "Simple Rail Fence Cipher"
	app.Usage = "简单栅栏加密解密器"
	app.Version = "1.0.0"
	app.Author = "One2r"
	app.Email = "601941036@qq.com"

	app.Commands = []cli.Command{
		{
			Name:      "encode",
			Usage:     "栅栏加密操作",
			Description:	"对明文进行加密操作", 
			Action: func(c *cli.Context) {
				if len(c.Args())<1 {
					fmt.Println("请输入需要进行加密操作的明文!")
					os.Exit(1)	
				}
				 
				rails := 2
				if c.Args().Get(1) != "" {
					r,err := strconv.Atoi(c.Args().Get(1))
					if err != nil {
						fmt.Println("参数："+c.Args().Get(1)+" 不能转换为整型！")
						os.Exit(1)
					}else{
						rails = r
					}
				}
				fmt.Println(rails)
			},
		},
		{
			Name:      "decode",
			Usage:     "栅栏解密操作",
			Description:	"对密文进行解密操作", 
			Action: func(c *cli.Context) {
				str := c.Args().First()
				fmt.Printf("密文为：%s\n",str)
				strlen := len(str)
				field:=[]int{}
				for i:=2;i<strlen;i++ {
					if strlen%i == 0 {
						field = append(field,i)
					}
				}
				fmt.Println("\n解密情况：")
				for i:=0;i<len(field);i++{
					line := strlen/field[i]
					result := ""
					for k:=0;k<field[i];k++ {
						for j:=0; j<line; j++{
							result += string(str[j*field[i]+k])
						}
					}
					fmt.Printf("%d) %d栏栅栏解密：\n",i+1,field[i])
					fmt.Printf("%s\n\n",result)
				}	
			},
		}}
	
	fmt.Printf("\n========%s========\n",app.Name)	
	app.Run(os.Args)
}
