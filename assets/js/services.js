/**
 * Created by serhii on 1/29/16.
 */

app.factory('Tools', function () {
    var tools = {};

    tools.cut = function (text, min) {
        if (min === undefined) {
            min = 200;
        }
        var newParagraph = text.indexOf("\n\n", min);
        if (newParagraph > -1) {
            return text.slice(0, newParagraph);
        }
        return text;
    };

    tools.date = function (ts) {
        return new Date(ts * 1000).format('Y-m-d H:i:s');
    };

    return tools;
});