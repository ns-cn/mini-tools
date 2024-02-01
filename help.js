
const COPY_MOD = "PASTE-MINI-COPY-MODEL"
const MUTE = "PASTE-MINI-MUTE"
function showHint(callbackSetList) {
    const copyAndKeep = utools.dbStorage.getItem(COPY_MOD)
    const mute = utools.dbStorage.getItem(MUTE)
    callbackSetList([
        {
            title: `一、🔧拷贝设置`,
            description: `当前模式: ${copyAndKeep ? "□拷贝但不隐藏窗口(适合简易编辑模式)" : "☑拷贝并隐藏窗口(适合工作模式)"},(选中切换)`,
            selectable: (itemData, callbackSetList) => {
                utools.dbStorage.setItem(COPY_MOD, !copyAndKeep)
                showHint(callbackSetList)
            }
        },
        {
            title: `二、⏰提醒设置`,
            description: `当前模式: ${mute ? "□拷贝时不提醒" : "☑拷贝时提醒(弹窗,需开启utools通知权限)"},(选中切换)`,
            selectable: (itemData, callbackSetList) => {
                utools.dbStorage.setItem(MUTE, !mute)
                showHint(callbackSetList)
            }
        },
        {
            title: `三、♻️清空所有已存储记录`,
            description: `清空所有已存储的剪贴板记录`,
            selectable: (itemData, callbackSetList) => {
                utools.dbStorage.removeItem("PASTE-MINI")
                utools.showNotification(`已清空所有剪贴板历史记录`)
            }
        },
        {
            title: `四、⌨️全局快捷键`,
            description: `utools-个人中心-全局快捷键: 设置快捷键,功能提示词设置为"剪贴板MINI"(选择拷贝功能提示词)`,
            selectable: (itemData, callbackSetList) => {
                utools.copyText("剪贴板MINI");
                utools.showNotification(`已拷贝"剪贴板MINI"`)
            }
        },
        { title: `五、🔍搜索模式`, description: `在搜索框中输入记录的文本或则自定义的tag即可过滤` },
        { title: `六、✍🏻编辑模式`, description: `选择带有✍🏻图标的记录即可进入编辑模式,在搜索框中输入记录的标签后,再次选择带有✍🏻的记录则退出编辑` },
        { title: `七、📒剪切板MINI介绍`, description: `插件提供剪切板中文本的复制、标记TAG等功能` },
        { title: `八、❓帮助（?）`, description: `输入问号?查看插件帮助信息` },
        { title: `九、📮联系作者`, description: `插件评论区捕捉野生插件作者，或则邮件至ns-cn@qq.com（注明utools插件）` },
    ])
}

module.exports = {
    showHint,
}
