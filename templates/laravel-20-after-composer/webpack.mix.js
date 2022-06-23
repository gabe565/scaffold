{{- .ExistingData }}

{{- if and (.NpmDeps.ModuleEnabled "browser-sync") (.NpmDeps.ModuleEnabled "browser-sync-webpack-plugin") }}

mix.browserSync({
    open: false,
    proxy: false,
    server: false,
    socket: { domain: 'clevyr.run:443' },
});
{{- end }}
