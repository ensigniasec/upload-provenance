# upload-provenance-action

Ensignia Provenance Action

Usage:

```yaml
steps:
  build:
    permissions:
      id-token: write # To sign the provenance.
      contents: write # To upload assets to release.
      actions: read # To read the workflow path.
      checks: read
    needs: args
    uses: slsa-framework/slsa-github-generator/.github/workflows/builder_go_slsa3.yml@v1.6.0
    with:
      go-version: 1.20.3

  publish:
    needs: build
    runs-on: ubuntu-latest
    steps:
      - name: Download artifact
        uses: actions/download-artifact@v3
        with:
          name: ${{ needs.build.outputs.go-provenance-name }}

      - uses: ensigniasec/upload-provenance-action@main
        with:
          api-key: "${{ secrets.ENSIGNIA_API_KEY }}"
          repo-token: ${{ secrets.GITHUB_TOKEN }}
          provenance-name: "${{ needs.build.outputs.go-provenance-name }}"
```
