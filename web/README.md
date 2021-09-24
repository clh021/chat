```bash
npx create-vite-app web

# fix lint when push web

cd web
npm install
npm install -D tailwindcss@npm:@tailwindcss/postcss7-compat @tailwindcss/postcss7-compat postcss@^7 autoprefixer@^9
npx tailwindcss init -p
```