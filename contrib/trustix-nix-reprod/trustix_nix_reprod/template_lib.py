from markupsafe import Markup
import urllib.parse
import json


def drv_url_quote(drv_path: str) -> str:
    return urllib.parse.quote(urllib.parse.quote(drv_path, safe=""))


def json_render(x) -> str:
    return json.dumps(x)


def js_url(filename: str) -> Markup:
    return Markup("/".join(("", "js", filename)))