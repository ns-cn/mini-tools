const pages = require("./pages")
const db = require("./core_db")
window.exports = {
    "paste": {
        mode: "list",
        args: {
            enter: (action, callbackSetList) => {
                db.watchMemeory(callbackSetList)
                window.currentPage = pages.LIST
                console.log(window.currentPage.type)
                window.currentPage.enter(action, callbackSetList)
            },
            search: (action, searchWord, callbackSetList) => {
                console.log("页面触发搜索:",window.currentPage.type)
                window.currentPage.search(action, searchWord, callbackSetList)
            },
            select: (action, itemData, callbackSetList) => {
                console.log("页面触发选择:",window.currentPage.type)
                window.currentPage.select(action, itemData, callbackSetList)
            },
            placeholder: "输入？打开帮助"
        }
    }
}

