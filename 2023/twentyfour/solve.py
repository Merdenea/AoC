import sympy

hs = [tuple(map(int, line.replace("@", ",").split(","))) for line in open("./in", "r")]

xr, yr, zr, vxr, vyr, vzr = sympy.symbols("xr, yr, zr, vxr, vyr, vzr")

equations = []

for sx, sy, sz, vx, vy, vz in hs:
    equations.append((xr - sx) * (vy - vyr) - (yr - sy) * (vx - vxr))
    equations.append((yr - sy) * (vz - vzr) - (zr - sz) * (vy - vyr))

res = sympy.solve(equations)
print(res[0][xr] +res[0][yr] + res[0][zr])

