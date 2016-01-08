# SimpleRailFenceCipher
简单栅栏加密解密器——Go语言版

Go版本：1.5  

Go包依赖：  
github.com/codegangsta/cli

#Features  
##加密
./SimpleRailFenceCipher encode "明文" [栅栏数]  
加密时在不能均匀分栏的情况下，将会在明文末尾补充“.”。  
###示例  
    ./SimpleRailFenceCipher encode "the anwser is wctf{C01umnar},if u is a big new,u can help us think more question,tks." 17

	========Simple Rail Fence Cipher========
	输入明文为：the anwser is wctf{C01umnar},if u is a big new,u can help us think more question,tks.
	加密栅栏数为：17
	
	加密结果：tn c0afsiwal kes,hwit1r  g,npt  ttessfu}ua u  hmqik e {m,  n huiouosarwCniibecesnren.

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
