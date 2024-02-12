
let positive = 0
let negative = 0

let testConnection = async function () {
  try {
    await fetch('https://fastly.com', { mode: 'no-cors' })
  } catch (e) {
    console.log('Error:', e)
    return e
  }
  return null
}

let loopTest = async function () {
  let element = document.getElementById('log')
  for (let i = 0; i < 20; i++) {
    let result = await testConnection()
    let text = `Test ${i}: `

    if (result == null) {
      text += 'Success'
      positive++
    } else {
      text += `Failed: ${result}`
      negative++
    }
    let li = document.createElement('li')
    li.innerText = text
    element.appendChild(li)
  }
}

window.onload = function () {
  loopTest().then(
    () => {
      document.getElementById('positive').innerText = positive
      document.getElementById('negative').innerText = negative
      let written = document.getElementById('written')
      if (negative > 0) {
        written.innerText = "Your network has problems"
        written.style.color = "red"
      } else {
        written.innerText = "Your network is fine"
        written.style.color = "green"
      }
    }
  )
}
