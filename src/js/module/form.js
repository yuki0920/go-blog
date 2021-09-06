'use strict';

// DOM Tree の構築が完了したら処理を開始
document.addEventListener('DOMContentLoaded', () => {
  // DOM API を利用して HTML 要素を取得
  const inputs = document.getElementsByTagName('input');
  const form = document.forms.namedItem('article-form');
  const saveBtn = document.querySelector('.article-form__save');
  const cancelBtn = document.querySelector('.article-form__cancel');
  const previewOpenBtn = document.querySelector('.article-form__open-preview');
  const previewCloseBtn = document.querySelector('.article-form__close-preview');
  const articleFormBody = document.querySelector('.article-form__body');
  const articleFormPreview = document.querySelector('.article-form__preview');
  const articleFormBodyTextArea = document.querySelector('.article-form__input--body');
  const articleFormPreviewTextArea = document.querySelector('.article-form__preview-body-contents');
  const errors = document.querySelector('.article-form__errors');
  const errorTmpl = document.querySelector('.article-form__error-tmpl').firstElementChild;

  // 新規作成画面か編集画面かを URL から判定
  const mode = { method: '', url: '' };
  if (window.location.pathname.endsWith('new')) {
    // 新規作成時の HTTP メソッドは POST を利用
    mode.method = 'POST';
    // 作成リクエスト、および戻るボタンの遷移先のパスは "/" になります。
    mode.url = '/';
  } else if (window.location.pathname.endsWith('edit')) {
    // 更新時の HTTP メソッドは PATCH を利用
    mode.method = 'PATCH';
    // 更新リクエスト、および戻るボタンの遷移先のパスは "/:articleID" になります。
    mode.url = `/${window.location.pathname.split('/')[1]}`;
  }
  const { method, url } = mode;

  // input 要素にフォーカスが合った状態で Enter が押されると form が送信されます。
  // 今回は Enter キーで form が送信されないように挙動を制御
  for (let elm of inputs) {
    elm.addEventListener('keydown', event => {
      if (event.keyCode && event.keyCode === 13) {
        // キーが押された際のデフォルトの挙動をキャンセル
        event.preventDefault();

        // 何もせず処理を終了
        return false;
      }
    });
  }

  // プレビューを開くイベントを設定
  previewOpenBtn.addEventListener('click', event => {
    // form の「本文」に入力された Markdown を HTML に変換してプレビューに埋め込みます。
    articleFormPreviewTextArea.innerHTML = md.render(articleFormBodyTextArea.value);
    // 入力フォームを非表示に
    articleFormBody.style.display = 'none';

    // プレビューを表示
    articleFormPreview.style.display = 'grid';
  });

  // プレビューを閉じるイベントを設定
  previewCloseBtn.addEventListener('click', event => {
    // 入力フォームを表示
    articleFormBody.style.display = 'grid';

    // プレビューを非表示に
    articleFormPreview.style.display = 'none';
  });

  // 前のページに戻るイベントを設定
  cancelBtn.addEventListener('click', event => {
    // <button> 要素クリック時のデフォルトの挙動をキャンセル
    event.preventDefault();

    // URL を指定して画面を遷移させます。
    window.location.href = url;
  });

  // 保存処理を実行するイベントを設定します。
  saveBtn.addEventListener('click', event => {
    event.preventDefault();

    errors.innerHTML = null;

    // フォームに入力された内容を取得します。
    const fd = new FormData(form);

    let status;

    // fetch API を利用してリクエストを送信します。
    fetch(url, {
      method: method,
      body: fd
    })
      .then(res => {
        status = res.status;
        return res.json();
      })
      .then(body => {
        console.log(JSON.stringify(body));

        if (status === 200) {
          // 成功時は一覧画面に遷移させます。
          window.location.href = url;
        }

        if (body.ValidationErrors) {
          showErrors(body.ValidationErrors);
        }
      })
      .catch(err => console.error(err));
  });

  const showErrors = messages => {
    if (Array.isArray(messages) && messages.length != 0) {
      const fragment = document.createDocumentFragment();

      messages.forEach(message => {
        const frag = document.createDocumentFragment();

        frag.appendChild(errorTmpl.cloneNode(true));

        frag.querySelector('.article-form__error').innerHTML = message;

        fragment.appendChild(frag);
      });

      errors.appendChild(fragment);
    }
  };
});
