# Solidity简介
    Solidity 是一种高级编程语言，由以太坊团队开发专门用于构建智能合约
# OpenZeppelin标准合约库
    一套由OpenZeppelin维护的标准智能合约库，能够加速合约的开发流程。
# Ethernaut
    一款游戏，通过攻击部署在游戏中的靶合约，学习关于合约安全的知识
# Hardhat与Truffle
    智能合约的测试与部署框架
# Web3.js
    一个前端JavaScript库，用于与智能合约进行交互。 在这份入门教程中，我们将引导你使用Solidity来进行基础的智能合约开发

# VS Code 配置Solidity开发环境
    1.新建一个空文件夹打开命令行，初始化node.js项目
        cd l1/solidity
        npm init -y    //文件夹中添加package.json
    2.使用npm下载hardhat（使用管理员）
        npm install --save-dev hardhat   //文件夹中添加node_modules，package-lock.json
    3.初始化项目
        创建hardhat项目
        npx hardhat init
    4.编写代码
    5.编译合约
        npx hardhat compile
    6.测试合约
        npx hardhat test
    7.部署合约
        在文件夹中 ignition/modules有一个文件名为Lock.js，用以下命令部署
        npx hardhat ignition deploy ./ignition/modules/Lock.js
        npx hardhat ignition deploy ./ignition/modules/Lock.js --network localhost
        或
        npx hardhat run ./ignition/modules/Lock.js
    8.链接钱包到hardhat network
        启动hardhat所带的链，它将公开一个json-rpc接口
        npx hardhat node
        注：如果端口被占用，使用
            netstat -ano 找到对应的pid
            taskkill /f -pid 160752
    9.hardhat-gas-reporter是hardhat的一个插件，专门用于查看每个单元测试所消耗的gas费用
    安装hardhat-gas-reporter
    ​npm install hardhat-gas-reporter --save-dev
    在配置hardhat-confing.js文件中添加
    ![](C:\Users\chenchen\AppData\Roaming\Typora\typora-user-images\image-20241124211148122.png)

    配置项

    ![image-20241124211311288](C:\Users\chenchen\AppData\Roaming\Typora\typora-user-images\image-20241124211311288.png)

    enable:表示使用report插件，修改为false则可以关闭

    currency：表示用人民币和gas费用进行换算，也可以设置为美元等

    ### 创建erc20 合约

    ### 部署合约到sepolia

    ### Links

    https://www.alchemy.com/

    https://learnblockchain.cn/article/8140

    https://faucet.quicknode.com/ethereum/sepolia

    https://blog.csdn.net/m0_73094011/article/details/135854126

    https://sepolia/etherscan.io/

    下载（管理员）

    npm i @openzeppelin/contracts

# 合约组成
    状态变量: 存储在合约中的变量，用于记录合约的持续状态。在本例中，定义了私有状态变量 owner 和 state。
    函数: 合约中定义的可执行代码块。本例包括构造函数 constructor 和其他函数 changeOwner、removeOwner、getOwner 以及 getState。
    函数修饰器: 用于修改函数行为的声明，这在函数定义前使用。在本例中，使用了 isOwner 修饰器以确保只有所有者可以执行某些操作。
    事件: 用于记录合约操作的日志，这有助于外部监听并追踪合约的活动。本例中引入了 OwnerSet 和 OwnerRemoved 事件。
    错误处理: 在合约中定义的错误处理机制，用于管理执行过程中的异常情况。
    结构体: 用于定义更复杂的数据类型。本例中创建了 Identity 结构体来存储额外信息。
    枚举: 用于定义一组命名常量，提高代码的可读性和维护性。本例中使用了 State 枚举来描述合约的不同状态。

# 数据类型
    值类型：这类类型的数据在传递时会被复制，每次传递都是一个独立的副本。
        布尔类型
        整型
        静态浮点型
        地址类型
        静态字节数组
        枚举类型
        自定义值类型
        字面值
    引用类型：相对于值类型，引用类型的数据在传递时不复制其本身，而是传递对原始数据的引用。
        数组
        结构体
        映射类型

# 数据位置
    Solidity 中有三种数据存储位置，分别指定变量的存储方式：
    storage： 数据永久存储在区块链上，通常用于状态变量。
    memory： 数据存储在内存中，是临时的，当函数调用结束后数据就会消失。
    calldata： 数据存储在一个专门用于存放函数参数的区域，这也是临时的。类似于 memory，但用于存放函数参数，与 memory 不同的是，calldata 中的数据不可修改且相比 memory 更节省 gas。 storage 可以类比为硬盘，而 memory 可类比为 RAM。calldata 可能稍显陌生，它的独特之处在于其不可变性和高效的 gas 使用。因其特性，当引用类型的函数参数不需要修改时，推荐使用 calldata 而非 memory。 为了避免过度复杂化，我们将在「Solidity进阶」中更深入地讨论 calldata 与 memory 的差异。目前，只需了解上述关于 calldata 的基本差异即可：仅用于函数参数，数据不可更改，是易失的，并且比 memory 更节约 gas。这些理解将帮助你更有效地使用Solidity中的数据位置