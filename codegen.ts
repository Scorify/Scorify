import type { CodegenConfig } from '@graphql-codegen/cli';
 
const config: CodegenConfig = {
  overwrite: true,
  schema: "src/graph/schema.graphqls",
  documents: "src/graph/schema.graphql",
  generates: {
    'src/graph/index.ts': {
      plugins: ['typescript', 'typescript-react-apollo', 'typescript-operations'],
    },
  },
};
export default config;