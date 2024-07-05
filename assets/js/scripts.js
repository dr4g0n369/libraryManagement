function createMessage(promptMessage, message) {
    const container = document.getElementById("Message");
    container.innerHTML = `<span class="close-btn" onclick="this.parentElement.classList.add('hidden-message')">&times;</span><strong>${promptMessage}</strong> ${message}`;
    container.classList.remove('hidden-message')
}

function createMessageUser(userid, username, role) {
    const container = document.getElementById("Message");
    container.innerHTML = `<span class="close-btn" onclick="this.parentElement.classList.add('hidden-message')">&times;</span><strong>UserId:</strong> ${userid} <br><strong>Username:</strong> ${username} <br><strong>Role:</strong> ${role}`;
    container.classList.remove('hidden-message')
}

function removeBook(bookid) {
    console.log("Removing a book...");
    let data = new FormData
    data.append("bookid", bookid)
    console.log(data.get("bookid"))
    fetch("/home/admin/removebook", {method: "POST", body: data})
        .then( res => {
            if (res.status == 200) {
                return res.json() 
            } else {
                return {success: -1}
            }
        })
        .then( data => {
            console.log(data)
            if (data.success == 1) {
                document.getElementById(`book-${data.bookid}`).innerHTML = ""
                createMessage("Success!", "Book removed successfully.")
            } else if (data.success = -1) {
                createMessage("Failed!", "Failed to remove book.")
            }
        })
    .catch(err => {
        console.log(err)
    })
}

function issueBook(bookid) {
    console.log("Returning a book...");
    let sendData = new FormData
    sendData.append("bookid", bookid)
    fetch("/home/issuebook", {method: "POST", body: sendData})
        .then( res => {
            if (res.status == 200) {
                return res.json() 
            } else {
                return {success: -1}
            }
        })
        .then( data => {
            if (data.success == 1) {
                document.getElementById(`book-${data.bookid}`).innerHTML = ""
                createMessage("Success!", "Book issued successfully.")
            } else if (data.success = -1) {
                createMessage("Failed!", "Failed to issue book.")
            }
        })
    .catch(err => {
        console.log(err)
    })
}

function returnBook(bookid) {
    console.log("Returning a book...");
    let data = new FormData
    data.append("bookid", bookid)
    fetch("/home/returnbook", {method: "POST", body: data})
        .then( res => {
            if (res.status == 200) {
                return res.json() 
            } else {
                return {success: -1}
            }
        })
        .then( data => {
            if (data.success == 1) {
                document.getElementById(`book-${data.bookid}`).innerHTML = ""
                createMessage("Success!", "Book returned successfully.")
            } else if (data.success = -1) {
                createMessage("Failed!", "Failed to return book.")
            }
        })
    .catch(err => {
        console.log(err)
    })
}

function getUserDetails(userid) {
    console.log("Returning a book...");
    let data = new FormData
    data.append("id", userid)
    fetch("/home/admin/getuserdetails", {method: "POST", body: data})
        .then( res => {
            if (res.status == 200) {
                return res.json() 
            } else {
                return {success: -1}
            }
        })
        .then( data => {
            if (data.success == 1) {
                createMessageUser(data.id, data.username, data.role)
            } else if (data.success = -1) {
                createMessage("Failed!", "Could not get user details.")
            }
        })
    .catch(err => {
        console.log(err)
    })
}
