# SimpleRailFenceCipher
简单栅栏加密解密器——Go语言版

Go版本：1.5  

Go包依赖：  
github.com/codegangsta/cli

#Features  
##加密
./SimpleRailFenceCipher encode "明文"

##解密
./SimpleRailFenceCipher decode "密文"
###示例
    ./SimpleRailFenceCipher decode "tn c0afsiwal kes,hwit1r  g,npt  ttessfu}ua u  hmqik e {m,  n huiouosarwCniibecesnren."

	========Simple Rail Fence Cipher========
	密文为：tn c0afsiwal kes,hwit1r  g,npt  ttessfu}ua u  hmqik e {m,  n huiouosarwCniibecesnren.
	
	解密情况：
	1) 5栏栅栏解密：
	taastg su km uwbnnfl,1, sah ,hoCer s hrntf me usncecikw ptuuq  iaien0wei te} i{noris.
	
	2) 17栏栅栏解密：
	the anwser is wctf{C01umnar},if u is a big new,u can help us think more question,tks.

#LICENSE
Licensed under the MIT license
