const { createApp } = Vue;

createApp({
    data() {
        return {
            // すべての行列（matrix）データをdata内で一括管理
            arena: Array.from({ length: 20 }, () => new Array(12).fill(0)),
            player: {
                pos: { x: 0, y: 0 },
                matrix: [[]],
            },
            score: 0,
            gameOver: false,
            dropCounter: 0,
            dropInterval: 1000,
            lastTime: 0,
            context: null,
            colors: [
                null, '#FF0D72', '#0DC2FF', '#0DFF72', 
                '#F538FF', '#FF8E0D', '#FFE138', '#3877FF'
            ]
        }
    },
    mounted() {
        this.context = this.$refs.canvas.getContext('2d');
        this.context.scale(20, 20);
        
        window.addEventListener('keydown', this.handleKeydown);
        this.resetGame();
        this.update();
    },
    methods: {
        createPiece(type) {
            const pieces = {
                'T': [[0, 1, 0], [1, 1, 1], [0, 0, 0]],
                'O': [[2, 2], [2, 2]],
                'L': [[0, 3, 0], [0, 3, 0], [0, 3, 3]],
                'J': [[0, 4, 0], [0, 4, 0], [4, 4, 0]],
                'I': [[0, 5, 0, 0], [0, 5, 0, 0], [0, 5, 0, 0], [0, 5, 0, 0]],
                'S': [[0, 6, 6], [6, 6, 0], [0, 0, 0]],
                'Z': [[7, 7, 0], [0, 7, 7], [0, 0, 0]],
            };
            return pieces[type];
        },
        draw() {
            this.context.fillStyle = '#000';
            this.context.fillRect(0, 0, this.$refs.canvas.width, this.$refs.canvas.height);
            
            // 盤面の描画
            this.drawMatrix(this.arena, { x: 0, y: 0 });
            // プレイヤーの描画
            this.drawMatrix(this.player.matrix, this.player.pos);
        },
        drawMatrix(matrix, offset) {
            matrix.forEach((row, y) => {
                row.forEach((value, x) => {
                    if (value !== 0) {
                        this.context.fillStyle = this.colors[value];
                        this.context.fillRect(x + offset.x, y + offset.y, 0.9, 0.9);
                    }
                });
            });
        },
        collide() {
            const [m, o] = [this.player.matrix, this.player.pos];
            for (let y = 0; y < m.length; ++y) {
                for (let x = 0; x < m[y].length; ++x) {
                    if (m[y][x] !== 0 && (this.arena[y + o.y] && this.arena[y + o.y][x + o.x]) !== 0) {
                        return true;
                    }
                }
            }
            return false;
        },
        merge() {
            this.player.matrix.forEach((row, y) => {
                row.forEach((value, x) => {
                    if (value !== 0) {
                        this.arena[y + this.player.pos.y][x + this.player.pos.x] = value;
                    }
                });
            });
        },
        playerRotate(dir) {
            const pos = this.player.pos.x;
            let offset = 1;
            this.rotateMatrix(this.player.matrix, dir);
            while (this.collide()) {
                this.player.pos.x += offset;
                offset = -(offset + (offset > 0 ? 1 : -1));
                if (offset > this.player.matrix[0].length) {
                    this.rotateMatrix(this.player.matrix, -dir);
                    this.player.pos.x = pos;
                    return;
                }
            }
        },
        rotateMatrix(matrix, dir) {
            for (let y = 0; y < matrix.length; ++y) {
                for (let x = 0; x < y; ++x) {
                    [matrix[x][y], matrix[y][x]] = [matrix[y][x], matrix[x][y]];
                }
            }
            dir > 0 ? matrix.forEach(row => row.reverse()) : matrix.reverse();
        },
        playerDrop() {
            if (this.gameOver) return;
            this.player.pos.y++;
            if (this.collide()) {
                this.player.pos.y--;
                this.merge();
                this.playerReset();
                this.arenaSweep();
            }
            this.dropCounter = 0;
        },
        playerMove(dir) {
            this.player.pos.x += dir;
            if (this.collide()) this.player.pos.x -= dir;
        },
        playerReset() {
            const pieces = 'ILJOTSZ';
            this.player.matrix = this.createPiece(pieces[pieces.length * Math.random() | 0]);
            this.player.pos.y = 0;
            this.player.pos.x = (this.arena[0].length / 2 | 0) - (this.player.matrix[0].length / 2 | 0);
            
            if (this.collide()) {
                this.gameOver = true;
            }
        },
        arenaSweep() {
            let rowCount = 1;
            for (let y = this.arena.length - 1; y > 0; --y) {
                if (this.arena[y].every(value => value !== 0)) {
                    const row = this.arena.splice(y, 1)[0].fill(0);
                    this.arena.unshift(row);
                    ++y;
                    this.score += rowCount * 10;
                    rowCount *= 2;
                }
            }
        },
        resetGame() {
            this.arena.forEach(row => row.fill(0));
            this.score = 0;
            this.gameOver = false;
            this.playerReset();
        },
        update(time = 0) {
            const deltaTime = time - this.lastTime;
            this.lastTime = time;
            this.dropCounter += deltaTime;
            
            if (this.dropCounter > this.dropInterval) {
                this.playerDrop();
            }
            
            this.draw();
            requestAnimationFrame(this.update);
        },
        handleKeydown(event) {
            if (this.gameOver) return;
            if (event.keyCode === 37) this.playerMove(-1); // Left
            else if (event.keyCode === 39) this.playerMove(1);  // Right
            else if (event.keyCode === 40) this.playerDrop();  // Down
            else if (event.keyCode === 81) this.playerRotate(-1); // Q
            else if (event.keyCode === 87) this.playerRotate(1);  // W
        }
    }
}).mount('#app');
