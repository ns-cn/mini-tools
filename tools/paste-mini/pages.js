const TYPE_LIST = "PAGE_LIST"
const TYPE_EDIT = "PAGE_EDIT"
const TYPE_HELP = "PAGE_HELP"
const db = require("./core_db")

const PageList = {
    type: TYPE_LIST,
    enter: (action, callbackSetList) => {
        utools.setSubInputValue("")
        callbackSetList(db.getNewest())
    },
    search: (action, searchWord, callbackSetList) => {
        if (searchWord === "?" || searchWord === "？") {
            PageHelp.enter(action, callbackSetList, PageList)
            window.currentPage = PageHelp
            // window.currentPage.enter(action, callbackSetList, PageList)
        } else if (window.editing) {
            window.eitingTag = searchWord
            if (!searchWord.startsWith(db.EDITING_TIPS)) {
                utools.setSubInputValue(db.EDITING_TIPS + searchWord)
            }
        } else {
            window.keyword = searchWord
            callbackSetList(db.getNewest())
        }
    },
    select: (action, itemData, callbackSetList) => {
        if (itemData.selectable) {
            itemData.selectable(itemData, callbackSetList)
        }
    },
}


const PageHelp = {
    type: TYPE_HELP,
    refresh: (action, callbackSetList) => {
        showHint(callbackSetList)
    },
    enter: (action, callbackSetList, sourcePage) => {
        showHint(callbackSetList)
    },
    search: (action, searchWord, callbackSetList) => {
        console.log("帮助页面搜索")
        if (searchWord !== "?" || searchWord !== "？") {
            window.currentPage = PageList
            window.currentPage.search(action, searchWord, callbackSetList)
        }
    },
    select: (action, itemData, callbackSetList) => {
        if (itemData.selectable) {
            itemData.selectable(itemData, callbackSetList)
        }
    },
}

function showHint(callbackSetList) {
    const copyAndKeep = db.getDbStorageItem(db.CACHE_COPY_MOD)
    const mute = db.getDbStorageItem(db.CACHE_MUTE)
    callbackSetList([
        {
            title: `一、🔧拷贝设置`,
            description: `当前模式: ${copyAndKeep ? "□拷贝但不隐藏窗口(适合简易编辑模式)" : "☑拷贝并隐藏窗口(适合工作模式)"},(选中切换)`,
            selectable: (itemData, callbackSetList) => {
                db.setDbStorageItem(db.CACHE_COPY_MOD, !copyAndKeep)
                showHint(callbackSetList)
            }
        },
        {
            title: `二、⏰提醒设置`,
            description: `当前模式: ${mute ? "□拷贝时不提醒" : "☑拷贝时提醒(弹窗,需开启utools通知权限)"},(选中切换)`,
            selectable: (itemData, callbackSetList) => {
                db.setDbStorageItem(db.CACHE_MUTE, !mute)
                showHint(callbackSetList)
            }
        },
        {
            title: `三、♻️清空所有已存储记录`,
            description: `清空所有已存储的剪贴板记录`,
            selectable: (itemData, callbackSetList) => {
                db.removeDbStorageItem(db.CACHE_ALL)
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
    LIST: PageList,
    TYPE_LIST,
    TYPE_EDIT,
    HELP: PageHelp,
    TYPE_HELP,
}
