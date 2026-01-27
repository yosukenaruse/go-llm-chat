const { createApp } = Vue;

createApp({
    data() {
        return {
            messages: [],
            newMessage: '',
            isLoading: false
        };
    },
    methods: {
        async sendMessage() {
            // 入力が空の場合は何もしない
            if (!this.newMessage.trim()) {
                return;
            }

            // ユーザーメッセージを追加
            const userMessage = {
                type: 'user',
                content: this.newMessage,
                timestamp: new Date()
            };
            this.messages.push(userMessage);

            // 送信するメッセージを保存して入力をクリア
            const messageToSend = this.newMessage;
            this.newMessage = '';

            // ローディング状態を設定
            this.isLoading = true;

            try {
                // Chat APIにPOSTリクエストを送信
                const response = await fetch('/chat', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        message: messageToSend
                    })
                });

                if (!response.ok) {
                    throw new Error('API request failed');
                }

                const data = await response.json();

                // アシスタントの返信を追加
                const assistantMessage = {
                    type: 'assistant',
                    content: data.reply,
                    timestamp: new Date()
                };
                this.messages.push(assistantMessage);

            } catch (error) {
                console.error('Error sending message:', error);
                
                // エラーメッセージを表示
                const errorMessage = {
                    type: 'assistant',
                    content: 'エラーが発生しました。もう一度お試しください。',
                    timestamp: new Date()
                };
                this.messages.push(errorMessage);
            } finally {
                this.isLoading = false;
                
                // メッセージ送信後、スクロールを最下部に移動
                this.$nextTick(() => {
                    this.scrollToBottom();
                });
            }
        },
        scrollToBottom() {
            const chatMessages = this.$refs.chatMessages;
            if (chatMessages) {
                chatMessages.scrollTop = chatMessages.scrollHeight;
            }
        }
    },
    mounted() {
        // 初期メッセージを表示（オプション）
        this.messages = [
            {
                type: 'assistant',
                content: 'こんにちは！何かお手伝いできることはありますか？',
                timestamp: new Date()
            }
        ];
    }
}).mount('#app');
