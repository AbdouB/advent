import re

DIGITS = 'zero one two three four five six seven eight nine'.split()
NUMBER_PAT = re.compile(f'(?=([0-9]|{"|".join(DIGITS)}))')


for line in open('input.txt').read().split('\n'):
    matches = NUMBER_PAT.findall(line)
    print(
        int(
            ''.join(
                str(DIGITS.index(m)) if m in DIGITS
                else m
                for m in [matches[0], matches[-1]]
            )
        )
    )

res = sum(
    int(
        ''.join(
            str(DIGITS.index(m)) if m in DIGITS
            else m
            for m in [matches[0], matches[-1]]
        )
    )
    for line in open('input.txt').read().split('\n')
    if (matches := NUMBER_PAT.findall(line))
)

print(res)