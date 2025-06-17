- name: Deploy to Render
run: |
    curl -fsSL https://render.com/deploy > deploy.sh  # Tải script deploy
    chmod +x deploy.sh
    ./deploy.sh ${{ secrets.RENDER_API_KEY }}  # Chạy script deploy và truyền API Key từ GitHub Secrets
env:
    RENDER_API_KEY: ${{ secrets.RENDER_API_KEY }}  # Lấy API Key từ GitHub Secrets