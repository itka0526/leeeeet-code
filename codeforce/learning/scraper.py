import os
import requests
import bs4
import re
import pdfkit


def get_result(path, file_name):
    file_name_without_ext = file_name.strip(".cpp")

    soup = request_problem(file_name_without_ext)
    markdown = soup_to_html(file_name_without_ext, soup)
    result = add_code(path + file_name, markdown)

    return result


def request_problem(problem_id):
    problem_url = f"https://acmp.ru/index.asp?main=task&id_task={problem_id}"
    response = requests.get(problem_url)
    response.encoding = response.apparent_encoding

    soup = bs4.BeautifulSoup(response.text, 'html.parser')

    soup = soup.find("td", {
        "background": "/images/notepad2.gif", "colspan": "3", "height": "100%", "style": "BORDER-BOTTOM: #306E18 1px dashed;  BORDER-left: #306E18 1px dashed;BORDER-right: #306E18 1px dashed; padding: 3px"
    })

    for tag in soup.find_all('h4'):
        tag.decompose()

    for tag in soup.find_all('script'):
        tag.decompose()

    for tag in soup.descendants:
        if tag.name == 'image' or tag.name == 'img':
            tag.attrs['src'] = 'https://acmp.ru' + tag.attrs['src']
        elif tag.name == 'table' or tag.name == 'td' or tag.name == 'th':
            tag.attrs = {
                'style': 'border: 1px solid black; border-collapse: collapse; padding: 2px 6px 2px 6px; '}
        elif isinstance(tag, bs4.element.Tag):
            tag.attrs = {}

    for element in soup(text=lambda text: isinstance(text, bs4.Comment)):
        element.extract()

    soup.attrs = {}
    soup.name = 'div'

    return soup


def soup_to_html(title, soup: bs4.Tag | bs4.NavigableString | None):
    tables = soup.find_all('table')
    cp = []
    for table in tables:
        cp.append(table.extract())

    title = '<h2 style="page-break-before: always;"> ЗАДАЧА №' + \
        str(int(title)) + "</h2>"

    str_html = soup.prettify()

    result = title + '\n' + str_html + '\n'

    for table in cp:
        result += '\n' + table.prettify()
    return result


def add_code(path, ct) -> str:
    f = open(path, "+r")
    code = f.read()
    f.close()

    code = code.replace("<", "&lt").replace(">", "&gt")

    ct += f"""
<pre><code class="prettyprint">{code}</code></pre>
"""

    return ct + '\n'


def valid_problem_str(s: str):
    example_str = '1346.cpp'
    return len(s) == len(example_str)


def ends_with_cpp(s: str):
    return s.endswith('.cpp')


def filter_invalid_problems(s):
    return ends_with_cpp(s) and valid_problem_str(s)


def save_to_html():
    curr_path = os.path.abspath(
        "learning/acmp/competitive-programming") + "/"
    res_file_name = "result.md"

    files = list(filter(filter_invalid_problems, os.listdir(curr_path)))
    files.sort()

    results = []

    for fn in files:
        print(f"{(files.index(fn) / len(files)) * 100.0:.2f}%. Current file: {fn}")
        res = get_result(curr_path, fn)

        if "не найдена!" in res:
            continue

        difficulty_procentage_idx = res.find(
            "Сложность: ") + len("Сложность: ")
        difficulty = int(re.findall(
            r'\d+', res[difficulty_procentage_idx: difficulty_procentage_idx + 3])[0])
        results.append((difficulty, res))

    results.sort(key=lambda item: item[0])

    final_result = '<script src="https://cdn.jsdelivr.net/gh/google/code-prettify@master/loader/run_prettify.js"></script>' + \
        '\n' + "\n".join([html for (_, html) in results])

    pdfkit.from_string(final_result, "acmp_ru_dmkz_solutions.pdf", options={
        'encoding': 'utf-8',
        'page-size': 'A4',
        'disable-smart-shrinking': '    '
    })

    return curr_path + res_file_name


save_to_html()
