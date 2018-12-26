extends Control

func _ready():
	_print_translated_message()

func _on_switch_to_nl_pressed():
	TranslationServer.set_locale('nl')
	_print_translated_message()


func _on_switch_to_en_pressed():
	TranslationServer.set_locale('en')
	_print_translated_message()

func _on_switch_to_no_pressed():
	TranslationServer.set_locale('no')
	_print_translated_message()

func _print_translated_message():
	print(TranslationServer.get_locale())
	print(TranslationServer.translate("TRANSLATED_PRINT_MSG"))

