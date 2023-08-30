import subprocess

import customtkinter as ctk

ctk.set_appearance_mode("dark")
ctk.set_default_color_theme("blue")


class AppGui(ctk.CTk):
    def __init__(self):
        super().__init__()
        self.title("Basic C App with Python GUI")
        self.geometry("500x350")

        self.frame = ctk.CTkFrame(master=self)
        self.frame.pack(fill="both", expand=True)

        self.label = ctk.CTkLabel(master=self.frame, text="Hello from Python!")
        self.label.pack()

        self.button = ctk.CTkButton(self, text="Call C main function", command=self.button_callbck)
        self.button.pack(padx=20, pady=20)

        self.output_label = ctk.CTkLabel(master=self.frame, text="")
        self.output_label.pack()

    def button_callbck(self):
        print("button clicked")
        try:
            result = subprocess.run(["./cmake-build-debug/GUI_BasicApplication"], stdout=subprocess.PIPE, text=True,
                                    check=True)
            output = result.stdout
            self.output_label.configure(text=output)
        except FileNotFoundError:
            print("Error: subprocess failed")


gui = AppGui()
gui.mainloop()