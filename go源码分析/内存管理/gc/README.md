
## GOGC

GOGC 是Go Runtime最早支持的环境变量，甚至比GOROOT还早，几乎无人不知。GOGC 用于控制GC的处发频率， 其值默认为100,

意为直到自上次垃圾回收后heap size已经增长了100%时GC才触发运行。即是GOGC=100意味着live heap size 每增长一倍，GC触发运行一次。

如设定GOGC=200, 则live heap size 自上次垃圾回收后，增长2倍时，GC触发运行， 总之，其值越大则GC触发运行频率越低， 反之则越高，

如果GOGC=off 则关闭GC.

虽然go 1.5引入了低延迟的GC, 但是GOGC对GC运行频率的影响不变， 仍然是其值大于100,则越大GC运行频率越高，

反之则越低。