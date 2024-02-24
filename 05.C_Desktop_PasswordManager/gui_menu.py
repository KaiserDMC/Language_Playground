import ctk as ctk
import os
from PIL import Image

ctk.set_appearance_mode("dark")
ctk.set_default_color_theme("blue")


class ScrollableLabelButtonFrame(ctk.CTkScrollableFrame):
    def __init__(self, master, command=None, **kwargs):
        super().__init__(master, **kwargs)
        self.grid_columnconfigure(0, weight=1)

        self.command = command
        self.radiobutton_variable = ctk.StringVar()
        self.label_list = []
        self.button_list = []

    def add_item(self, item, image=None):
        label = ctk.CTkLabel(self, text=item, image=image, compound="left", padx=5, anchor="w")
        button = ctk.CTkButton(self, text="Copy Password", width=100, height=24)
        if self.command is not None:
            button.configure(command=lambda: self.command(item))
        label.grid(row=len(self.label_list), column=0, pady=(0, 10), sticky="w")
        button.grid(row=len(self.button_list), column=1, pady=(0, 10), padx=5)
        self.label_list.append(label)
        self.button_list.append(button)

    def remove_item(self, item):
        for label, button in zip(self.label_list, self.button_list):
            if item == label.cget("text"):
                label.destroy()
                button.destroy()
                self.label_list.remove(label)
                self.button_list.remove(button)
                return


class Menu(ctk.CTk):
    width = 700
    height = 450

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.title("Password Manager")
        self.geometry(f"{self.width}x{self.height}")
        self.resizable(False, False)

        # set grid layout 1x2
        self.grid_rowconfigure(0, weight=1)
        self.grid_columnconfigure(1, weight=1)

        # password check for list
        self.valid_password_provided = True

        # load images with light and dark mode image
        image_path = os.path.join(os.path.dirname(os.path.realpath(__file__)), "images")
        self.logo_image = ctk.CTkImage(Image.open(os.path.join(image_path, "lock-solid.png")), size=(26, 26))
        self.large_test_image = ctk.CTkImage(Image.open(os.path.join(image_path, "test-banner.png")), size=(500, 150))
        # self.image_icon_image = ctk.CTkImage(Image.open(os.path.join(image_path, "image_icon_light.png")), size=(20, 20))
        # self.home_image = ctk.CTkImage(light_image=Image.open(os.path.join(image_path, "home_dark.png")),
        #                                        dark_image=Image.open(os.path.join(image_path, "home_light.png")), size=(20, 20))
        self.frame_3_image = ctk.CTkImage(light_image=Image.open(os.path.join(image_path, "vault-solid.png")),
                                          dark_image=Image.open(os.path.join(image_path, "vault-solid.png")),
                                          size=(20, 20))
        # self.add_user_image = ctk.CTkImage(light_image=Image.open(os.path.join(image_path, "add_user_dark.png")),
        # dark_image=Image.open(os.path.join(image_path, "add_user_light.png")), size=(20, 20))

        # create navigation frame
        self.navigation_frame = ctk.CTkFrame(self, corner_radius=0)
        self.navigation_frame.grid(row=0, column=0, sticky="nsew")
        self.navigation_frame.grid_rowconfigure(5, weight=1)

        self.navigation_frame_label = ctk.CTkLabel(self.navigation_frame, text="  Image Example", image=self.logo_image,
                                                   compound="left", font=ctk.CTkFont(size=15, weight="bold"))
        self.navigation_frame_label.grid(row=0, column=0, padx=20, pady=20)

        self.home_button = ctk.CTkButton(self.navigation_frame, corner_radius=0, height=40, border_spacing=10,
                                         text="Home",
                                         fg_color="transparent", text_color=("gray10", "gray90"),
                                         hover_color=("gray70", "gray30"),
                                         anchor="w", command=self.home_button_event)
        self.home_button.grid(row=1, column=0, sticky="ew")

        self.frame_2_button = ctk.CTkButton(self.navigation_frame, corner_radius=0, height=40, border_spacing=10,
                                            text="Generate Credentials",
                                            fg_color="transparent", text_color=("gray10", "gray90"),
                                            hover_color=("gray70", "gray30"),
                                            anchor="w", command=self.frame_2_button_event)
        self.frame_2_button.grid(row=2, column=0, sticky="ew")

        self.frame_3_button = ctk.CTkButton(self.navigation_frame, corner_radius=0, height=40, border_spacing=10,
                                            text="Stored Passwords",
                                            fg_color="transparent", text_color=("gray10", "gray90"),
                                            hover_color=("gray70", "gray30"), image=self.frame_3_image,
                                            anchor="w", command=self.frame_3_button_event)
        self.frame_3_button.grid(row=3, column=0, sticky="ew")

        self.about_button = ctk.CTkButton(self.navigation_frame, corner_radius=0, height=40, border_spacing=10,
                                          text="About",
                                          fg_color="transparent", text_color=("gray10", "gray90"),
                                          hover_color=("gray70", "gray30"),
                                          anchor="w", command=self.about_frame_button_event)
        self.about_button.grid(row=4, column=0, sticky="ew")

        self.appearance_mode_menu = ctk.CTkOptionMenu(self.navigation_frame, values=["Dark", "Light", "System"],
                                                      command=self.change_appearance_mode_event)
        self.appearance_mode_menu.grid(row=6, column=0, padx=20, pady=20, sticky="s")

        # create home frame
        self.home_frame = ctk.CTkFrame(self, corner_radius=0, fg_color="transparent")
        self.home_frame.grid_columnconfigure(0, weight=1)

        self.home_frame_large_image_label = ctk.CTkLabel(self.home_frame, text="", image=self.large_test_image)
        self.home_frame_large_image_label.grid(row=0, column=0, padx=20, pady=50)

        self.home_frame_button_1 = ctk.CTkButton(self.home_frame, text="Create new creds file",
                                                 command=self.popup_window_event, width=30)
        self.home_frame_button_1.grid(row=1, column=0, padx=20, pady=(50, 20), sticky="nsew", columnspan=2)
        self.home_frame_button_1._min_height = 60

        self.home_frame_button_2 = ctk.CTkButton(self.home_frame, text="Load existing creds file",
                                                 compound="right", width=30)
        self.home_frame_button_2.grid(row=2, column=0, padx=20, pady=(20, 10), sticky="nsew", columnspan=2)
        self.home_frame_button_2._min_height = 60

        # create second frame
        self.second_frame = ctk.CTkFrame(self, corner_radius=0, fg_color="transparent")
        self.second_frame.grid_columnconfigure(0, weight=1)

        label_1 = ctk.CTkLabel(self.second_frame, text="Website/Application:")
        label_1.grid(row=0, column=0, padx=20, pady=(10, 0), sticky="w")
        entry_1 = ctk.CTkEntry(self.second_frame)
        entry_1.grid(row=1, column=0, padx=20, pady=(0, 10), sticky="ew")

        label_2 = ctk.CTkLabel(self.second_frame, text="Username:")
        label_2.grid(row=2, column=0, padx=20, pady=(10, 0), sticky="w")
        entry_2 = ctk.CTkEntry(self.second_frame)
        entry_2.grid(row=3, column=0, padx=20, pady=(0, 10), sticky="ew")

        label_3 = ctk.CTkLabel(self.second_frame, text="Password:")
        label_3.grid(row=4, column=0, padx=20, pady=(10, 0), sticky="w")
        entry_3 = ctk.CTkEntry(self.second_frame)
        entry_3.grid(row=5, column=0, padx=20, pady=(0, 10), sticky="ew")

        button_left = ctk.CTkButton(self.second_frame, text="Generate Password")
        button_left.grid(row=6, column=0, padx=20, pady=(10, 0), sticky="w")

        button_right = ctk.CTkButton(self.second_frame, text="Save Credentials")
        button_right.grid(row=6, column=0, padx=20, pady=(10, 0), sticky="e")

        # create third frame
        self.third_frame = ctk.CTkFrame(self, corner_radius=0, fg_color="transparent")
        current_dir = os.path.dirname(os.path.abspath(__file__))

        self.third_frame_button = ScrollableLabelButtonFrame(master=self.third_frame, width=490, height=440,
                                                             command=self.label_button_frame_event,
                                                             corner_radius=0)

        self.third_frame_button.grid(row=2, column=0, padx=20, pady=10)
        if self.valid_password_provided:
            for i in range(20):  # add items with images
                self.third_frame_button.add_item(f"image and item {i}")

        # create about frame
        self.about_frame = ctk.CTkFrame(self, corner_radius=0, fg_color="transparent")
        self.about_frame.grid_columnconfigure(0, weight=1)
        self.about_frame.grid_columnconfigure(1, weight=1)
        self.about_frame_label = ctk.CTkLabel(self.about_frame, text="About",
                                              font=ctk.CTkFont(size=15, weight="bold"))
        self.about_frame_label.grid(row=0, column=0, columnspan=2, padx=10, pady=10, sticky="nsew")

        about_text = ("This is software is under the BSD 3-Clause License.\n\n Copyright (c) 2023, Otaku Devs Team\n\n "
                      "Ver 0.0.1 by Krisztian Fodor")
        about_label = ctk.CTkLabel(self.about_frame, text=about_text, anchor="w")
        about_label.grid(row=1, column=1, padx=10, pady=125, sticky="nsew")

        self.about_frame.grid(row=0, column=2, columnspan=2, padx=0, pady=0, sticky="nsew")

        # select default frame
        self.select_frame_by_name("home")

    def select_frame_by_name(self, name):
        # set button color for selected button
        self.home_button.configure(fg_color=("gray75", "gray25") if name == "home" else "transparent")
        self.frame_2_button.configure(fg_color=("gray75", "gray25") if name == "frame_2" else "transparent")
        self.frame_3_button.configure(fg_color=("gray75", "gray25") if name == "frame_3" else "transparent")
        self.about_button.configure(fg_color=("gray75", "gray25") if name == "about" else "transparent")

        # show selected frame
        if name == "home":
            self.home_frame.grid(row=0, column=1, sticky="nsew")
        else:
            self.home_frame.grid_forget()
        if name == "frame_2":
            self.second_frame.grid(row=0, column=1, sticky="nsew")
        else:
            self.second_frame.grid_forget()
        if name == "frame_3":
            self.third_frame.grid(row=0, column=1, sticky="nsew")
        else:
            self.third_frame.grid_forget()
        if name == "about":
            self.about_frame.grid(row=0, column=1, sticky="nsew")
        else:
            self.about_frame.grid_forget()

    def home_button_event(self):
        self.select_frame_by_name("home")

    def frame_2_button_event(self):
        self.select_frame_by_name("frame_2")

    def frame_3_button_event(self):
        self.select_frame_by_name("frame_3")

    def about_frame_button_event(self):
        self.select_frame_by_name("about")

    def change_appearance_mode_event(self, new_appearance_mode):
        ctk.set_appearance_mode(new_appearance_mode)

    def label_button_frame_event(self, item):
        print(f"label button frame clicked: {item}")

    def popup_window_event(self):
        result = ctk.CTkInputDialog(text="Enter master password", title="Master Password")
        if result:
            # Here, result will contain the value entered by the user.
            # You can use this value as needed, e.g., for verifying the master password.
            print("Entered Master Password:", result)


app = Menu()
app.mainloop()
