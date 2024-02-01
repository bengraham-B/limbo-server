const testAPI  = async () => {
    const response = await fetch("http://localhost:8080/create", {
        method: "POST",
        body: JSON.stringify({
            fileName: "py.py",
            fileType: "text",
            fileContent: "46, 334, 2267, 145, 13, 12, 56, 34, 23, 78, 52, 576, 67,8"
        }),
        headers: {
            "Content-Type":"application/json"
        }
    })

    const data = await response.json()

    console.log(data)
}

testAPI()