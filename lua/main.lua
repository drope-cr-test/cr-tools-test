-- C-T4: references a global declared ONLY in .luacheckrc.
-- If the rc file executes, no "undefined global" finding.
-- If the rc is ignored, luacheck flags CR_MARKER_T4_X.
print(CR_MARKER_T4_X)
print(CR_OTHER_UNDECLARED)

-- C-T round 2: valid Lua comment for the test change
