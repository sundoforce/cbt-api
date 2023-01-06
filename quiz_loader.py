import csv

def text_to_csv(filename):
    rows = []
    with open(filename, 'r', encoding='utf-8') as f:
        lines = f.read().strip().split('\n')
        for line in lines:
            print(line)
            if line and line[0] == '*':
                row = line[2:].split('. ')
                rows.append(row)
    with open(filename.replace('.txt', '.csv'), 'w', newline='', encoding='utf-8') as csvfile:
        writer = csv.writer(csvfile, delimiter=',')
        for row in rows:
            writer.writerow(row)

# 예시
filename = "quiz.txt"
text_to_csv(filename)
