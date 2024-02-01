const { clipboard, ipcMain } = require('electron')
const COPY_MOD = "PASTE-MINI-COPY-MODEL"
const MUTE = "PASTE-MINI-MUTE"
const EDITING_TIPS = "输入标签:"
window.exports = {
    "paste": {
        mode: "list",
        args: {
            enter: (action, callbackSetList) => {
                let timer = null
                if (timer) {
                    clearInterval(timer)
                }
                let last = ""
                timer = setInterval(() => {
                    // 读取剪贴板内容  
                    const text = clipboard.readText()
                    if (last !== text && text !== "") {
                        last = text
                        update(text)
                        if (!window.editing) {
                            callbackSetList(getNewest())
                        }
                    }
                }, 500)
                callbackSetList(getNewest())
            },
            search: (action, searchWord, callbackSetList) => {
                if (searchWord === "?" || searchWord === "？") {
                    require('./help').showHint(callbackSetList);
                } else if (window.editing) {
                    window.eitingTag = searchWord
                    if (!searchWord.startsWith(EDITING_TIPS)) {
                        utools.setSubInputValue(EDITING_TIPS + searchWord)
                    }
                } else {
                    window.keyword = searchWord
                    callbackSetList(getNewest())
                }
            },
            select: (action, itemData, callbackSetList) => {
                if (itemData.selectable) {
                    itemData.selectable(itemData, callbackSetList)
                }
            },
            placeholder: "输入？打开帮助"
        }
    }
}

function display(item, clipboardText) {
    const editable = item.data === clipboardText
    let title = item.data
    if (editable) {
        title = `✍🏻${item.data}`
    }
    return {
        ...item,
        title: title,
        editable: editable,
        description: `更新时间: ${item.time}| 标签: ${item.tag}`,
        selectable: (itemData, callbackSetList) => {
            if (!itemData.editable) {
                const copyAndKeep = utools.dbStorage.getItem(COPY_MOD)
                const mute = utools.dbStorage.getItem(MUTE)
                if (copyAndKeep) {
                    utools.copyText(itemData.data);
                } else {
                    utools.hideMainWindowPasteText(itemData.data)
                }
                if (!mute) {
                    utools.showNotification(`已拷贝${itemData.title}`)
                }
            } else {
                if (window.editing) {
                    update(itemData.data, window.eitingTag.substring(5).trim())
                    callbackSetList(getNewest())
                    utools.setSubInputValue(window.keyword ? window.keyword : "")
                    window.editing = false
                } else {
                    window.editing = true
                    utools.setSubInputValue(`${EDITING_TIPS}${itemData.tag}`)
                }
            }
        }
    }
}

function getNewest() {
    let target = utools.dbStorage.getItem("PASTE-MINI")
    const clipboardText = clipboard.readText()
    // 将window.keyword按照空格进行分割
    const words = window.keyword ? window.keyword.split(" ").filter(t => t !== "") : []
    if (target) {
        if (window.keyword && window.keyword !== "") {
            return target.filter(item => {
                // 必须words中的所有关键词都命中才返回true
                if (words.length === 0) {
                    return true
                }
                if (words.every(word => {
                    return item.data?.includes(word) || item.tag?.includes(word)
                })) {
                    return true
                }
                return false
            }).sort((a, b) => {
                return b.time - a.time
            }).map(item => display(item, clipboardText))
        }
        return target.sort((a, b) => {
            return b.time - a.time
        }).map(item => display(item, clipboardText))
    }
    return []
}

function update(text, tag) {
    if (!text) {
        return
    }
    if (!tag) {
        tag = ""
    }
    let target = utools.dbStorage.getItem("PASTE-MINI")
    if (!target) {
        target = []
    }
    // 过滤text中data为text的
    target = target.filter(item => {
        return item.data !== text
    })
    // 在最前面增加一个最新的记录
    target.unshift({
        data: text,
        time: new Date(),
        tag: tag
    })
    // 仅保留target的前1000条
    target = target.slice(0, 1000)
    utools.dbStorage.setItem("PASTE-MINI", target)
}