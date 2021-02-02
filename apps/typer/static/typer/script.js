function init() {
    var request = new XMLHttpRequest();
    request.onreadystatechange = function () {
        if (request.readyState == 4 && request.status == "200") {
            var quotes = JSON.parse(request.responseText);
            startf(quotes.quotes[Math.floor(Math.random() * quotes.quotes.length)]);
        }
    }
    request.open("GET", "/static/typer/quotes.json");
    request.send();
}

function differ(text, quote) {
    html = "";
    green = "";
    i = 0;

    for (x in text) {
        if (text[x] == quote[x]) {
            green += quote[x];
            i++;
            if (text == quote) {
                state = 0;
            }
        } else {
            break;
        }
    }
    count = green.split(" ").length
    $("#count").text(count + " words");
    return html = '<span id="green">' + green + '</span><span id="red">' + quote.slice(i, text.length) + '</span>' + quote.slice(text.length, quote.length);
}

function wpmf() {
    count = $("#count").text();
    wpm = parseInt(count) / (end - start) * 60000;
    return Math.floor(wpm);
}

function typer() {
    state = 1;
    start = new Date;

    $("#textbox").removeAttr('disabled');
    $("#textbox").select();
    $("#textbox").keyup(function () {
        $("#quotebox").html(differ($("#textbox").val(), $("#quotebox").text()));
        if (state == 0) {
            end = new Date;
            let speed = wpmf()
            $("#count").text($("#count").text() + " --- " + ((end - start) / 1000) + "s --- " + speed + "wpm");
            $("#textbox").attr('disabled', true);

            $("#speed").val(speed)
            $("#submit").attr('hidden', false);
        }
    })
}

function counter(count) {
    $("#counter").show();
    if (count > 0) {
        $("#counter").text("00:0" + (count));
        setTimeout(counter, 1000, count - 1);
    } else {
        $("#counter").hide();
        return;
    }
}

function startf(quote) {
    $("#textbox").val('');
    $("#textbox").attr('disabled', true);
    $("#count").text('');
    counter(1);
    $("#quotebox").text(quote.quote);
    setTimeout(typer, 1000);
}

$(document).ready(function () {
    $(document).keypress(function (event) {
        if (event.which == 13) {
            init();
        }
    });
});
