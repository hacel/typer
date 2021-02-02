

function pickquote() {
    quotes = ["asdasd asd", "fdasd"]
    return quotes[Math.floor(Math.random() * quotes.length)]
}

function differ() {
    var textbox = document.getElementById('textbox');
    var quotebox = document.getElementById('quotebox');
    a = quotebox.textContent;
    b = textbox.value;
    // ret = '<span style="color: yellow;">' + quotebox.innerHTML + '</span>';
    // console.log(ret)

    // for (i = 0; i <= b.length; i++) {
    //     if (a[i] == b[i]) {

    //     }
    // }
    if (a == b) {
        console.log("E")
        document.getElementById("form").submit();
    } else {
        console.log("N", a, b)
    }
    // quotebox.innerHTML = ret
}


function typer() {
    textbox.removeAttribute('readonly');
    textbox.select();
}

function start() {
    var textbox = document.getElementById('textbox');
    var quotebox = document.getElementById('quotebox');

    quote = pickquote();
    quotebox.textContent = quote;

    setTimeout(typer, 2000);
}