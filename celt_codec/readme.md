# c文件来自 celt-0.7.1的libcelt 文件夹
+ 去掉了 `dump_modes.c`
+ 去掉了 `plc.c`
+ 去掉了 `testcelt.c`
+ 注释掉了 `celt.c` 部分代码,使其不调用到plc.c上去：
  ```
  /* #define NEW_PLC */
  #if !defined(FIXED_POINT) || defined(NEW_PLC)
  #include "plc.c"
  #endif
  ```