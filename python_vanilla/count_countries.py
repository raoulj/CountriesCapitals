from geopy.distance import geodesic
import pandas as pd

l = []

records = pd.read_csv('../country-capitals.csv').to_records(index=False)
for i in range(len(records)):
    for j in range(len(records)):
        if i >= j:
            continue
        a_country, a_capital, a_lat, a_long, _, _ = records[i]
        b_country, b_capital, b_lat, b_long, _, _ = records[j]
        l.append((f'{a_capital}, {a_country}', f'{a_capital}, {a_country}', geodesic((a_lat, a_long), (b_lat, b_long)).miles))
print('\n'.join([str(x) for x in sorted(l, key=lambda x: x[2])[:50]]))