// Copyright (C) 2017  Nicholas Anderssohn
import 'dart:html';
import 'package:simple_streams/simple_streams.dart';
import 'package:codemirror/codemirror.dart';
import 'dropdown/dropdown.dart';
import 'page/page.dart';

const String startingCCode = """#include "stdio.h"

int main() {
printf("Hello, C!\\n");
return 0;
}""";

const String startingCPPCode = """#include <iostream>

int main() {
  std::cout << "Hello, C++!" << std::endl;
  return 0;
}""";

const String startingPython3Code = """print("Hello, Python 3!")""";

const String startingGoCode = """package main

import "fmt"

func main() {
  fmt.Println("Hello, Go!")
}""";

const String startingJavaCode = """// Do not change the name of this class
public class Main {

    public static void main(String[] args) {
        System.out.println("Hello, Java!");
    }

}""";

class CodeEditor {
  SimpleStream _sendStream = new SimpleStream();
  CodeMirror _editor;
  DivElement target = new DivElement();
  TextAreaElement textArea;
  DivElement runBtn = new DivElement();
  DivElement output = new DivElement();
  DropDown dropDown;
  DivElement dropBtn = new DivElement();
  StandardBtn selectedLang;
  List langBtns = [
    new StandardBtn(new DivElement(), tags: ['main.c', startingCCode], text: "C"),
    new StandardBtn(new DivElement(), tags: ['main.cpp', startingCPPCode], text: "C++"),
    new StandardBtn(new DivElement(), tags: ['main.py', startingPython3Code], text: "Python 3"),
    new StandardBtn(new DivElement(), tags: ['main.go', startingGoCode], text: "Go"),
    new StandardBtn(new DivElement(), tags: ['Main.java', startingJavaCode], text: "Java")
  ];
  // bool waitingForOutput = false;
  String get language => dropBtn.text;
         set language(String value) {
           dropBtn.text = value;
           dropBtn.children.add(dropDown.target);
         }
  String get code => _editor.getDoc().getValue();
  set code(String value) {_editor.getDoc().setValue(value);}
  Map options = {
    'lineNumbers': true,
    'theme': 'neat',
    'continueComments': {'continueLineComment': false},
    'autoCloseTags': true,
    'mode': 'clike',
    'matchBrackets': true,
    'extraKeys': {
      'Ctrl-Space': 'autocomplete',
      'Cmd-/': 'toggleComment',
      'Ctrl-/': 'toggleComment' }
    };
  CodeEditor(String elementId, {options: null}) {
    selectedLang = langBtns[0];
    textArea = querySelector(elementId);
    runBtn..classes.add('send-code-btn')..text = "Run"..contentEditable = "false";
    textArea.classes.add('code-editor-text-area');
    target.children..add(runBtn)..add(output);
    if (options != null) this.options = options;
    _editor = new CodeMirror.fromTextArea(textArea, options: this.options);
    output..text = "output"..classes.add('output');

    runBtn.onClick.listen((var e) {
      _sendStream.add([selectedLang.tags[0], code]); // [filename, code]
    });

    target.children.add(dropBtn);
    dropBtn.classes.add('language-selector-btn');
    dropBtn.text = 'C';
    dropDown = new DropDown(dropBtn);
    dropDown.onClick((var e) => selectedLang.tags[1] = code);
    _initializeDropDown();
    code = selectedLang.tags[1];
  }

  _initializeDropDown() {
    for (var btn in langBtns) {
      btn.onClick(_handleLangBtnClicked);
      dropDown.addItem(btn.target);
    }
  }

  onSendClicked(handler(var e)) {
    _sendStream.listen(handler);
  }

  _handleLangBtnClicked(var e) {
    language = e.text;
    selectedLang = e;
    code = e.tags[1];
  }
}
