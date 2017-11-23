import plistlib
import json
import pdb
import pprint

result = []

pl = plistlib.readPlist("NationalCodeList.plist")
pl = sorted(pl, key=lambda k: k["country"])


for elem in pl:
    print({"name": elem.country, "code": elem.code})
    # result.append({"name": elem.country, "code": elem.code})
