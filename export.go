package main

/*
#include <lua.h>
#include <stdlib.h>

int invoke_go_func(lua_State* state) {
  void* p = lua_touserdata(state, lua_upvalueindex(1));
  return Invoke(p);
}

void register_function(lua_State* state, const char* name, void* func) {
  lua_pushlightuserdata(state, func);
  lua_pushcclosure(state, (lua_CFunction)invoke_go_func, 1);
  lua_setglobal(state, name);
}

int traceback(lua_State* L) {
  lua_getfield(L, LUA_GLOBALSINDEX, "debug");
  lua_getfield(L, -1, "traceback");
  lua_pushvalue(L, 1);
  lua_pushinteger(L, 2);
  lua_call(L, 2, 1);
  return 1;
}

void setup_message_handler(lua_State* L) {
  lua_pushcfunction(L, traceback);
}

*/
import "C"
