/* 
迪米特法则：如果两个类不必彼此直接通信，那么这两个类就不必直接发生作用。
如果一个类要调用另一类中的某一方法时，可以通过第三者转发这个调用。


迪米特法则其根本思想，是强调了类之间的松耦合。
就拿你今天碰到的这件事来做例子，你第一天去公司，怎么会认识IT部的人呢，如果公司有很好的管理，
那么应该是人事的小杨打个电话到IT部，告诉主管安排人给小菜你装电脑，就算开始是小张负责，
他临时有急事，主管也可以再安排小李来处理。同样道理，我们在程序设计时，类之间的耦合越弱，
越有利于复用，一个处在弱耦合的类被修改，不会对有关系的类造成波及。也就是说，信息的隐藏促进了软件的复用。

程杰. 大话设计模式 (Chinese Edition) (Kindle 位置 1851-1855). 清华大学出版社. Kindle 版本. 
*/