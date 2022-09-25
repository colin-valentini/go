import json

json_file = open('practice/algoexpert/problems.json')
data = json.load(json_file)
json_file.close()

difficulties = [1, 2, 3, 4]
difficulty_names = {
    1: "Easy",
    2: "Medium",
    3: "Hard",
    4: "Very Hard",
}
problems = {
    1: [],
    2: [],
    3: [],
    4: [],
}

questions = data['questions']
for q in questions:
    url = f"https://www.algoexpert.io/questions/{q['uid']}"
    line = f"- [ ] [{q['name']}]({url})\n"
    problems[q['difficulty']].append({
        "line": line,
        "name": q['name'],
    })

out_file = open('practice/algoexpert/algoexpert.md', 'w')

for difficulty in difficulties:
    out_file.write(f"### {difficulty_names[difficulty]}\n")
    diff_problems = problems[difficulty]
    diff_problems.sort(key=lambda p: p['name'])
    lines = [p['line'] for p in diff_problems]
    out_file.writelines(lines)
    out_file.write('\n')

out_file.close()

