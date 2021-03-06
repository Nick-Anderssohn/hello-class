// Copyright (C) 2017  Nicholas Anderssohn

import 'page.dart';
import '../../button/button.dart';
import '../../textbox/textbox.dart';
import 'create_comm_handler.dart';
import 'login_comm_handler.dart';
import 'dart:html';
import 'dart:async';
import 'package:simple_streams/simple_streams.dart';

class SelectorPage extends Page {
  String htmlDoc = '../resources/selector.html';
  SimpleStreamRouter _joinBackArrowRouter;
  SimpleStreamRouter _createBackArrowRouter;
  StandardBtn _selectCreatePageBtn;
  StandardBtn _selectJoinPageBtn;
  StandardBtn _educatorLoginBtn;
  StandardBtn _studentLoginBtn;
  StandardBtn _createClassBtn;
  TextBox _joinClassTextbox;
  PasswordBox _joinClassPasswordTextbox;
  TextBox _createClassTextbox;
  PasswordBox _createPasswordTextbox;
  CreateCommHandler _createCommHandler;
  LoginCommHandler _loginCommHandler;
  DivElement _createPage;
  DivElement _joinPage;
  DivElement _selectPage;
  DivElement _curPage;
  final String _left = "calc(-100% - 10px)";
  final String _middle = "0px";
  final String _right = "calc(100% + 10px)";
  static const int _animDuration = 900; // milliseconds

  SelectorPage();

  load() {
    if (!hasBeenLoaded) {
      _queryElementsFromHtml();
      _setupJoinPage();
      _setupCreatePage();
      hasBeenLoaded = true;
      _joinBackArrowRouter.listen(repositionPages);
      _createBackArrowRouter.listen(repositionPages);
    }
  }

  _queryElementsFromHtml() {
    _createPage = querySelector('#create-page');
    _joinPage = querySelector('#join-page');
    _curPage = _selectPage = querySelector('#select-page');
    _selectCreatePageBtn = new StandardBtn(querySelector('#select-create-page-btn'));
    _selectJoinPageBtn = new StandardBtn(querySelector('#select-join-page-btn'));
    _educatorLoginBtn = new StandardBtn(querySelector('#educator-login-btn'));
    _studentLoginBtn = new StandardBtn(querySelector('#student-login-btn'));
    _joinClassTextbox = new TextBox(querySelector('#join-class-textbox'), defaultText: 'Class Name');
    _joinClassPasswordTextbox = new PasswordBox(querySelector('#join-class-password-textbox'), defaultText: 'Password (Leave empty if student)');
    _createClassBtn = new StandardBtn(querySelector('#create-class-btn'));
    _createClassTextbox = new TextBox(querySelector('#create-class-textbox'), defaultText: 'Class Name');
    _createPasswordTextbox = new PasswordBox(querySelector('#create-password-textbox'), defaultText: 'Password');
    _joinBackArrowRouter = new SimpleStreamRouter(querySelector('#join-back-arrow').onClick);
    _createBackArrowRouter = new SimpleStreamRouter(querySelector('#create-back-arrow').onClick);
  }

  _setupJoinPage() {
    _selectJoinPageBtn.onClick(_selectJoinPageBtnClick);
    _loginCommHandler = new LoginCommHandler(_joinClassTextbox, _joinClassPasswordTextbox, _educatorLoginBtn, _studentLoginBtn, '/');
  }

  _setupCreatePage() {
    _selectCreatePageBtn.onClick(_selectCreatePageBtnClick);
    _createCommHandler = new CreateCommHandler(_createClassBtn, _createClassTextbox, _createPasswordTextbox, '/');
  }

  _selectJoinPageBtnClick(var e) {
    _slideOutLeft(_selectPage);
    _slideIn(_joinPage);
    _curPage = _joinPage;
  }

  _selectCreatePageBtnClick(var e) {
    _slideOutRight(_selectPage);
    _slideIn(_createPage);
    _curPage = _createPage;
  }

  _slideOutLeft(DivElement page) async {
    page.classes.add('page-out-left');
    new Future.delayed(const Duration(milliseconds: _animDuration + 100), () {
      page.style.left = _left;
      page.classes.remove('page-out-left');
    });
  }

  _slideOutRight(DivElement page) async {
    page.classes.add('page-out-right');
    new Future.delayed(const Duration(milliseconds: _animDuration + 100), () {
      page.style.left = _right;
      page.classes.remove('page-out-right');
    });
  }

  _slideIn(DivElement page) async {
    page.classes.add('page-in');
    new Future.delayed(const Duration(milliseconds: _animDuration + 100), () {
      page.style.left = _middle;
      page.classes.remove('page-in');
    });
  }

  repositionPages(var e) {
    if (_curPage == _selectPage) {
      return;
    }
    if (_curPage == _joinPage) {
      _slideOutRight(_joinPage);
    } else {
      _slideOutLeft(_createPage);
    }
    _slideIn(_selectPage);
  }
}
