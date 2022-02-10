const dupObjList = [
    {id: 1, score: 10},
    {id: 1, score: 3},
    {id: 1, score: 7},
    {id: 2, score: 7},
    {id: 2, score: 3},
    {id: 3, score: 3},
]

const parseObjToList = (obj) => {
    if (Object.keys(obj).length === 0) {
        return
    }

    let newList = []

    for (const [k, v] of Object.entries(obj)) {
        newList.push({
            id: parseInt(k, 10),
            score: v
        })
    }
    return newList
}

const removeDup = () => {
    const list = {}
    for (const x of dupObjList) {
        if (list[x.id]) {
            continue
        }
        list[x.id] = x.score
    }

    console.log('removeDup ans:', parseObjToList(list))
}


const lookup = () => {
    const list = {}
    for (const x of dupObjList) {
        if (list[x.id]) {
            list[x.id] += x.score
            continue
        }
        list[x.id] = x.score
    }

    console.log('lookup ans:', parseObjToList(list))
}

// removeDup()
// lookup()

const sortList = [4, 3, 1, 2]

const sortAscListNumber = () => {
    let tryCount = 0
}

// sortAscListNumber()

// const a = "cde"
// const b = "abc"
//
// const a = "showman"
// const b = "woman"

const a = "fcrxzwscanmligyxyvym"
const b = "jxwtrhvujlmrpdoqbisbwhmgpmeoke"

// expect 2
// showman
// woman



const stringAnagram = (x, y) => {
    const listX = [...x]
    const listY = [...y]

    const maxLength = listX.length > listY.length ? listX.length : listY.length

    let ans = []
    for (let i = 0; i < listX.length; i++) {
        for (let j = 0; j < listY.length; j++) {
            if (listX[i] === listY[j]) {
                ans.push(listX[i])
            }
        }
    }

    // console.log('stringAnagram ans:', (xLength + yLength) - 2)
    // console.log('ans',maxLength -  ans.length)
    console.log('ans',  ans)
}

stringAnagram(a, b)