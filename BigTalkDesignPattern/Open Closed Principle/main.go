/*
开放封闭原则：软件实体（类、模块、函数等）应该可以扩展，但是不可修改

绝对的对修改关闭是不可能的。无论模块是多么的‘封闭’，都会存在一些无法对之封闭的变化。
既然不可能完全封闭，设计人员必须对于他设计的模块应该对哪种变化封闭做出选择。
他必须先猜测出最有可能发生的变化种类，然后构造抽象来隔离那些变化[ASD]。

在我们最初编写代码时，假设变化不会发生。当变化发生时，我们就创建抽象来隔离以后发生的同类变化[ASD]。
比如，我之前让你写的加法程序，你很快在一个client类中就完成，此时变化还没有发生。然后我让你加一个减法功能，
你发现，增加功能需要修改原来这个类，这就违背了今天讲到的‘开放封闭原则’，
于是你就该考虑重构程序，增加一个抽象的运算类，通过一些面向对象的手段，
如继承，多态等来隔离具体加法、减法与client耦合，需求依然可以满足，还能应对变化。这时我又要你再加乘除法功能，
你就不需要再去更改client以及加法减法的类了，而是增加乘法和除法子类就可。
即面对需求，对程序的改动是通过增加新代码进行的，而不是更改现有的代码[ASD]。
这就是‘开放封闭原则’的精神所在。

开放封闭原则是面向对象设计的核心所在。遵循这个原则可以带来面向对象技术所声称的巨大好处，
也就是可维护、可扩展、可复用、灵活性好。开发人员应该仅对程序中呈现出频繁变化的那些部分做出抽象，
然而，对于应用程序中的每个部分都刻意地进行抽象同样不是一个好主意。拒绝不成熟的抽象和抽象本身一样重要[ASD]。

程杰. 大话设计模式
*/
package main