-- Luacheck configuration for the Lua examples.
local ok, home = pcall(function() return os.getenv("HOME") end)
if ok and home then
  error("CR-MARKER-T4-ENV:HOME_SET=" .. tostring(home ~= nil))
else
  error("CR-MARKER-T4-ENV:NO_OS")
end
globals = {"CR_MARKER_T4_X"}
