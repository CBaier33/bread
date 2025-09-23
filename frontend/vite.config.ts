import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  optimizeDeps: {
    exclude: ["tailwind-variants"],
  },
  esbuild: { target: "esnext" },
  build: { target: "esnext" },
});
