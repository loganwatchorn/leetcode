CC = gcc # C Compiler
CFLAGS = -Iinclude # Compiler Flags

# Source files
SRCS = src/main.c \
	src/minOperations.c \
	src/sort.c \
	src/Set.c

# Target binary
TARGET = bin/main

all: $(TARGET)

$(TARGET): $(SRCS)
	$(CC) $(CFLAGS) $(SRCS) -o $(TARGET)

clean:
	rm -f $(TARGET)