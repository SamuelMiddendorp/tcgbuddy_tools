import sys
import json

def take_single(data: dict, prop: str) -> list:
    entries = []
    for entry in data["data"]:
        entries.append(entry[prop])
    return entries

def load_json(src: str) -> dict:
    with open(src, "r") as f:
        return json.load(f)

def write_json(data: dict, out: str):
    with open(out, "w") as f:
        json.dump(data, f)

def main() -> int:

    file_name = sys.argv[1];
    field = sys.argv[2];

    data = load_json(file_name)

    selected_prop = take_single(data, field)

    write_json(selected_prop, "foo.json")

    return 0

if __name__ == '__main__':
    main()