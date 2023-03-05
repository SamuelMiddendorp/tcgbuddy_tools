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

    file_name = sys.argv[1] if len(sys.argv) >1 else "in.json";
    field = sys.argv[2] if len(sys.argv) >2 else "id";
    out_file_name = sys.argv[3] if len(sys.argv) >3 else "out.json"
    try:
        data = load_json(file_name)
    except:
        print("Unable to open file: {}".format(file_name))
        return 1
    try:
        selected_prop = take_single(data, field)
    except:
        print("Cannot find property: {}".format(field))
        return 1

    write_json(selected_prop, out_file_name)

    return 0

if __name__ == '__main__':
    main()