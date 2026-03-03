#ifndef MAL_ODIN_EDITLINE_H
#define MAL_ODIN_EDITLINE_H

char *readline(const char *prompt);
int add_history(const char *line);
int read_history(const char *filename);
int write_history(const char *filename);

#endif

