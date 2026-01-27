const { createApp } = Vue;

createApp({
    data() {
        return {
            password: '',
            errorMessage: '',
            successMessage: '',
            isLoading: false,
            showError: false,
            showSuccess: false
        };
    },
    methods: {
        async login() {
            // 入力が空の場合は何もしない
            if (!this.password.trim()) {
                return;
            }

            // メッセージをリセット
            this.showError = false;
            this.showSuccess = false;
            this.errorMessage = '';
            this.successMessage = '';

            // ローディング状態を設定
            this.isLoading = true;

            try {
                // Login APIにPOSTリクエストを送信
                const response = await fetch('/login', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        password: this.password
                    })
                });

                const data = await response.json();

                if (response.ok && data.success) {
                    // ログイン成功
                    this.successMessage = 'ログイン成功！リダイレクト中...';
                    this.showSuccess = true;
                    
                    // 1.5秒後にチャット画面にリダイレクト
                    setTimeout(() => {
                        window.location.href = '/static/chat.html';
                    }, 1500);
                } else {
                    // ログイン失敗
                    this.errorMessage = data.message || 'パスワードが正しくありません';
                    this.showError = true;
                    
                    // パスワード入力欄をクリア
                    this.password = '';
                    
                    // ローディング状態を解除
                    this.isLoading = false;
                    
                    // 入力欄にフォーカス
                    this.$nextTick(() => {
                        this.$refs.passwordInput.focus();
                    });
                }
            } catch (error) {
                console.error('Login error:', error);
                this.errorMessage = 'ログインに失敗しました。もう一度お試しください。';
                this.showError = true;
                
                // パスワード入力欄をクリア
                this.password = '';
                
                // ローディング状態を解除
                this.isLoading = false;
                
                // 入力欄にフォーカス
                this.$nextTick(() => {
                    this.$refs.passwordInput.focus();
                });
            }
        }
    },
    computed: {
        buttonText() {
            return this.isLoading ? 'ログイン中...' : 'ログイン';
        },
        isButtonDisabled() {
            return this.isLoading || !this.password.trim();
        }
    }
}).mount('#app');
