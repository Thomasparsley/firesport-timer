import tkinter as tk


def a():
    print("a")


def main():
    root = tk.Tk()
    root.title("Firesport timer")

    my_label = tk.Label(root, text="Hello World")
    my_label.pack()

    my_button = tk.Button(root, text="Click me", command=a)
    my_button.pack()

    root.mainloop()

    super_loop()


def super_loop():
    pass


if __name__ == '__main__':
    main()
