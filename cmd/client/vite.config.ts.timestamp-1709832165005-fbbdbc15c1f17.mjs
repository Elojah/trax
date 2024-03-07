// vite.config.ts
import { fileURLToPath, URL } from "node:url";
import { defineConfig } from "file:///home/elojah/go/src/github.com/elojah/trax/cmd/client/node_modules/vite/dist/node/index.js";
import vue from "file:///home/elojah/go/src/github.com/elojah/trax/cmd/client/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import vueJsx from "file:///home/elojah/go/src/github.com/elojah/trax/cmd/client/node_modules/@vitejs/plugin-vue-jsx/dist/index.mjs";
import fs from "fs";
var __vite_injected_original_import_meta_url = "file:///home/elojah/go/src/github.com/elojah/trax/cmd/client/vite.config.ts";
var vite_config_default = async () => {
  let config = {};
  const loadConfigPlugin = {
    name: "load-config",
    buildStart() {
      const configFile = process.env.CONFIG_FILE;
      if (!configFile) {
        throw new Error("CONFIG_FILE not set");
      }
      const configData = fs.readFileSync(configFile, "utf-8");
      if (!configData) {
        throw new Error("CONFIG_FILE is empty");
      }
      config = JSON.parse(configData);
    },
    transform(code) {
      return code.replace("__CONFIG__", JSON.stringify(config));
    }
  };
  return defineConfig({
    plugins: [
      vue(),
      vueJsx(),
      loadConfigPlugin
    ],
    resolve: {
      alias: {
        "@": fileURLToPath(new URL("./src", __vite_injected_original_import_meta_url)),
        "@internal": fileURLToPath(new URL("./../../internal", __vite_injected_original_import_meta_url)),
        "@pkg": fileURLToPath(new URL("./../../pkg", __vite_injected_original_import_meta_url)),
        "@api": fileURLToPath(new URL("./../../cmd/api/grpc", __vite_injected_original_import_meta_url)),
        // imports from external files
        "@protobuf-ts": fileURLToPath(new URL("./node_modules/@protobuf-ts", __vite_injected_original_import_meta_url))
      }
    }
  });
};
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCIvaG9tZS9lbG9qYWgvZ28vc3JjL2dpdGh1Yi5jb20vZWxvamFoL3RyYXgvY21kL2NsaWVudFwiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9maWxlbmFtZSA9IFwiL2hvbWUvZWxvamFoL2dvL3NyYy9naXRodWIuY29tL2Vsb2phaC90cmF4L2NtZC9jbGllbnQvdml0ZS5jb25maWcudHNcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfaW1wb3J0X21ldGFfdXJsID0gXCJmaWxlOi8vL2hvbWUvZWxvamFoL2dvL3NyYy9naXRodWIuY29tL2Vsb2phaC90cmF4L2NtZC9jbGllbnQvdml0ZS5jb25maWcudHNcIjtpbXBvcnQgeyBmaWxlVVJMVG9QYXRoLCBVUkwgfSBmcm9tICdub2RlOnVybCdcblxuaW1wb3J0IHsgZGVmaW5lQ29uZmlnIH0gZnJvbSAndml0ZSdcbmltcG9ydCB2dWUgZnJvbSAnQHZpdGVqcy9wbHVnaW4tdnVlJ1xuaW1wb3J0IHZ1ZUpzeCBmcm9tICdAdml0ZWpzL3BsdWdpbi12dWUtanN4J1xuaW1wb3J0IGZzIGZyb20gJ2ZzJztcblxuZXhwb3J0IGRlZmF1bHQgYXN5bmMoKSA9PntcbiAgbGV0IGNvbmZpZyA9IHt9XG4gIGNvbnN0IGxvYWRDb25maWdQbHVnaW4gPSB7XG4gICAgbmFtZTogJ2xvYWQtY29uZmlnJyxcbiAgICBidWlsZFN0YXJ0KCkge1xuICAgICAgY29uc3QgY29uZmlnRmlsZSA9IHByb2Nlc3MuZW52LkNPTkZJR19GSUxFO1xuXG4gICAgICBpZiAoIWNvbmZpZ0ZpbGUpIHtcbiAgICAgICAgdGhyb3cgbmV3IEVycm9yKCdDT05GSUdfRklMRSBub3Qgc2V0Jyk7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IGNvbmZpZ0RhdGEgPSBmcy5yZWFkRmlsZVN5bmMoY29uZmlnRmlsZSwgJ3V0Zi04Jyk7XG5cbiAgICAgIGlmICghY29uZmlnRGF0YSkge1xuICAgICAgICB0aHJvdyBuZXcgRXJyb3IoJ0NPTkZJR19GSUxFIGlzIGVtcHR5Jyk7XG4gICAgICB9XG5cbiAgICAgIGNvbmZpZyA9IEpTT04ucGFyc2UoY29uZmlnRGF0YSk7XG4gICAgfSxcbiAgICB0cmFuc2Zvcm0oY29kZTphbnkpIHtcbiAgICAgIHJldHVybiBjb2RlLnJlcGxhY2UoJ19fQ09ORklHX18nLCBKU09OLnN0cmluZ2lmeShjb25maWcpKTtcbiAgICB9XG4gIH1cblxuICByZXR1cm4gZGVmaW5lQ29uZmlnKHtcbiAgICBwbHVnaW5zOiBbXG4gICAgICB2dWUoKSxcbiAgICAgIHZ1ZUpzeCgpLFxuICAgICAgbG9hZENvbmZpZ1BsdWdpbixcbiAgICBdLFxuICAgIHJlc29sdmU6IHtcbiAgICAgIGFsaWFzOiB7XG4gICAgICAgICdAJzogZmlsZVVSTFRvUGF0aChuZXcgVVJMKCcuL3NyYycsIGltcG9ydC5tZXRhLnVybCkpLFxuICAgICAgICAnQGludGVybmFsJzogZmlsZVVSTFRvUGF0aChuZXcgVVJMKCcuLy4uLy4uL2ludGVybmFsJywgaW1wb3J0Lm1ldGEudXJsKSksXG4gICAgICAgICdAcGtnJzogZmlsZVVSTFRvUGF0aChuZXcgVVJMKCcuLy4uLy4uL3BrZycsIGltcG9ydC5tZXRhLnVybCkpLFxuICAgICAgICAnQGFwaSc6IGZpbGVVUkxUb1BhdGgobmV3IFVSTCgnLi8uLi8uLi9jbWQvYXBpL2dycGMnLCBpbXBvcnQubWV0YS51cmwpKSxcbiAgICAgICAgLy8gaW1wb3J0cyBmcm9tIGV4dGVybmFsIGZpbGVzXG4gICAgICAgICdAcHJvdG9idWYtdHMnOiBmaWxlVVJMVG9QYXRoKG5ldyBVUkwoJy4vbm9kZV9tb2R1bGVzL0Bwcm90b2J1Zi10cycsIGltcG9ydC5tZXRhLnVybCkpLFxuICAgICAgfVxuICAgIH1cbiAgfSlcbn1cbiJdLAogICJtYXBwaW5ncyI6ICI7QUFBaVYsU0FBUyxlQUFlLFdBQVc7QUFFcFgsU0FBUyxvQkFBb0I7QUFDN0IsT0FBTyxTQUFTO0FBQ2hCLE9BQU8sWUFBWTtBQUNuQixPQUFPLFFBQVE7QUFMbU0sSUFBTSwyQ0FBMkM7QUFPblEsSUFBTyxzQkFBUSxZQUFVO0FBQ3ZCLE1BQUksU0FBUyxDQUFDO0FBQ2QsUUFBTSxtQkFBbUI7QUFBQSxJQUN2QixNQUFNO0FBQUEsSUFDTixhQUFhO0FBQ1gsWUFBTSxhQUFhLFFBQVEsSUFBSTtBQUUvQixVQUFJLENBQUMsWUFBWTtBQUNmLGNBQU0sSUFBSSxNQUFNLHFCQUFxQjtBQUFBLE1BQ3ZDO0FBRUEsWUFBTSxhQUFhLEdBQUcsYUFBYSxZQUFZLE9BQU87QUFFdEQsVUFBSSxDQUFDLFlBQVk7QUFDZixjQUFNLElBQUksTUFBTSxzQkFBc0I7QUFBQSxNQUN4QztBQUVBLGVBQVMsS0FBSyxNQUFNLFVBQVU7QUFBQSxJQUNoQztBQUFBLElBQ0EsVUFBVSxNQUFVO0FBQ2xCLGFBQU8sS0FBSyxRQUFRLGNBQWMsS0FBSyxVQUFVLE1BQU0sQ0FBQztBQUFBLElBQzFEO0FBQUEsRUFDRjtBQUVBLFNBQU8sYUFBYTtBQUFBLElBQ2xCLFNBQVM7QUFBQSxNQUNQLElBQUk7QUFBQSxNQUNKLE9BQU87QUFBQSxNQUNQO0FBQUEsSUFDRjtBQUFBLElBQ0EsU0FBUztBQUFBLE1BQ1AsT0FBTztBQUFBLFFBQ0wsS0FBSyxjQUFjLElBQUksSUFBSSxTQUFTLHdDQUFlLENBQUM7QUFBQSxRQUNwRCxhQUFhLGNBQWMsSUFBSSxJQUFJLG9CQUFvQix3Q0FBZSxDQUFDO0FBQUEsUUFDdkUsUUFBUSxjQUFjLElBQUksSUFBSSxlQUFlLHdDQUFlLENBQUM7QUFBQSxRQUM3RCxRQUFRLGNBQWMsSUFBSSxJQUFJLHdCQUF3Qix3Q0FBZSxDQUFDO0FBQUE7QUFBQSxRQUV0RSxnQkFBZ0IsY0FBYyxJQUFJLElBQUksK0JBQStCLHdDQUFlLENBQUM7QUFBQSxNQUN2RjtBQUFBLElBQ0Y7QUFBQSxFQUNGLENBQUM7QUFDSDsiLAogICJuYW1lcyI6IFtdCn0K
