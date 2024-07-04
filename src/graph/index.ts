import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  JSON: { input: any; output: any; }
  Time: { input: any; output: any; }
  Upload: { input: any; output: any; }
};

export type Check = {
  __typename?: 'Check';
  config: Scalars['JSON']['output'];
  configs: Array<CheckConfig>;
  create_time: Scalars['Time']['output'];
  editable_fields: Array<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  source: Source;
  statuses: Array<Status>;
  update_time: Scalars['Time']['output'];
  weight: Scalars['Int']['output'];
};

export type CheckConfig = {
  __typename?: 'CheckConfig';
  check: Check;
  check_id: Scalars['ID']['output'];
  config: Scalars['JSON']['output'];
  create_time: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  update_time: Scalars['Time']['output'];
  user: User;
  user_id: Scalars['ID']['output'];
};

export type Config = {
  __typename?: 'Config';
  check: Check;
  config: Scalars['JSON']['output'];
  id: Scalars['ID']['output'];
  user: User;
};

export enum EngineState {
  Paused = 'paused',
  Running = 'running',
  Stopping = 'stopping',
  Waiting = 'waiting'
}

export type File = {
  __typename?: 'File';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  url: Scalars['String']['output'];
};

export type Inject = {
  __typename?: 'Inject';
  create_time: Scalars['Time']['output'];
  end_time: Scalars['Time']['output'];
  files: Array<File>;
  id: Scalars['ID']['output'];
  rubric: RubricTemplate;
  start_time: Scalars['Time']['output'];
  submissions: Array<InjectSubmission>;
  title: Scalars['String']['output'];
  update_time: Scalars['Time']['output'];
};

export type InjectSubmission = {
  __typename?: 'InjectSubmission';
  create_time: Scalars['Time']['output'];
  files: Array<File>;
  graded: Scalars['Boolean']['output'];
  id: Scalars['ID']['output'];
  inject: Inject;
  inject_id: Scalars['ID']['output'];
  notes: Scalars['String']['output'];
  rubric?: Maybe<Rubric>;
  update_time: Scalars['Time']['output'];
  user: User;
  user_id: Scalars['ID']['output'];
};

export type InjectSubmissionByUser = {
  __typename?: 'InjectSubmissionByUser';
  submissions: Array<InjectSubmission>;
  user: User;
};

export type LoginOutput = {
  __typename?: 'LoginOutput';
  domain: Scalars['String']['output'];
  expires: Scalars['Int']['output'];
  httpOnly: Scalars['Boolean']['output'];
  name: Scalars['String']['output'];
  path: Scalars['String']['output'];
  secure: Scalars['Boolean']['output'];
  token: Scalars['String']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  adminBecome: LoginOutput;
  adminLogin: LoginOutput;
  changePassword: Scalars['Boolean']['output'];
  createCheck: Check;
  createInject: Inject;
  createUser: User;
  deleteCheck: Scalars['Boolean']['output'];
  deleteInject: Scalars['Boolean']['output'];
  deleteUser: Scalars['Boolean']['output'];
  editConfig: Config;
  gradeSubmission: InjectSubmission;
  login: LoginOutput;
  sendGlobalNotification: Scalars['Boolean']['output'];
  startEngine: Scalars['Boolean']['output'];
  stopEngine: Scalars['Boolean']['output'];
  submitInject: InjectSubmission;
  updateCheck: Check;
  updateInject: Inject;
  updateUser: User;
};


export type MutationAdminBecomeArgs = {
  id: Scalars['ID']['input'];
};


export type MutationAdminLoginArgs = {
  id: Scalars['ID']['input'];
};


export type MutationChangePasswordArgs = {
  newPassword: Scalars['String']['input'];
  oldPassword: Scalars['String']['input'];
};


export type MutationCreateCheckArgs = {
  config: Scalars['JSON']['input'];
  editable_fields: Array<Scalars['String']['input']>;
  name: Scalars['String']['input'];
  source: Scalars['String']['input'];
  weight: Scalars['Int']['input'];
};


export type MutationCreateInjectArgs = {
  end_time: Scalars['Time']['input'];
  files: Array<Scalars['Upload']['input']>;
  rubric: RubricTemplateInput;
  start_time: Scalars['Time']['input'];
  title: Scalars['String']['input'];
};


export type MutationCreateUserArgs = {
  number?: InputMaybe<Scalars['Int']['input']>;
  password: Scalars['String']['input'];
  role: Role;
  username: Scalars['String']['input'];
};


export type MutationDeleteCheckArgs = {
  id: Scalars['ID']['input'];
};


export type MutationDeleteInjectArgs = {
  id: Scalars['ID']['input'];
};


export type MutationDeleteUserArgs = {
  id: Scalars['ID']['input'];
};


export type MutationEditConfigArgs = {
  config: Scalars['JSON']['input'];
  id: Scalars['ID']['input'];
};


export type MutationGradeSubmissionArgs = {
  rubric: RubricInput;
  submissionID: Scalars['ID']['input'];
};


export type MutationLoginArgs = {
  password: Scalars['String']['input'];
  username: Scalars['String']['input'];
};


export type MutationSendGlobalNotificationArgs = {
  message: Scalars['String']['input'];
  type: NotificationType;
};


export type MutationSubmitInjectArgs = {
  files: Array<Scalars['Upload']['input']>;
  injectID: Scalars['ID']['input'];
  notes: Scalars['String']['input'];
};


export type MutationUpdateCheckArgs = {
  config?: InputMaybe<Scalars['JSON']['input']>;
  editable_fields?: InputMaybe<Array<Scalars['String']['input']>>;
  id: Scalars['ID']['input'];
  name?: InputMaybe<Scalars['String']['input']>;
  weight?: InputMaybe<Scalars['Int']['input']>;
};


export type MutationUpdateInjectArgs = {
  add_files?: InputMaybe<Array<Scalars['Upload']['input']>>;
  delete_files?: InputMaybe<Array<Scalars['ID']['input']>>;
  end_time?: InputMaybe<Scalars['Time']['input']>;
  id: Scalars['ID']['input'];
  rubric?: InputMaybe<RubricTemplateInput>;
  start_time?: InputMaybe<Scalars['Time']['input']>;
  title?: InputMaybe<Scalars['String']['input']>;
};


export type MutationUpdateUserArgs = {
  id: Scalars['ID']['input'];
  number?: InputMaybe<Scalars['Int']['input']>;
  password?: InputMaybe<Scalars['String']['input']>;
  username?: InputMaybe<Scalars['String']['input']>;
};

export type Notification = {
  __typename?: 'Notification';
  message: Scalars['String']['output'];
  type: NotificationType;
};

export enum NotificationType {
  Default = 'default',
  Error = 'error',
  Info = 'info',
  Success = 'success',
  Warning = 'warning'
}

export type Query = {
  __typename?: 'Query';
  check: Check;
  checks: Array<Check>;
  config: Config;
  configs: Array<Config>;
  inject: Inject;
  injectSubmission: InjectSubmission;
  injectSubmissions: Array<InjectSubmission>;
  injectSubmissionsByUser: Array<InjectSubmissionByUser>;
  injects: Array<Inject>;
  me?: Maybe<User>;
  scoreboard: Scoreboard;
  source: Source;
  sources: Array<Source>;
  users: Array<User>;
};


export type QueryCheckArgs = {
  id?: InputMaybe<Scalars['ID']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
};


export type QueryConfigArgs = {
  id: Scalars['ID']['input'];
};


export type QueryInjectArgs = {
  id: Scalars['ID']['input'];
};


export type QueryInjectSubmissionArgs = {
  id: Scalars['ID']['input'];
};


export type QueryInjectSubmissionsByUserArgs = {
  id: Scalars['ID']['input'];
};


export type QueryScoreboardArgs = {
  round?: InputMaybe<Scalars['Int']['input']>;
};


export type QuerySourceArgs = {
  name: Scalars['String']['input'];
};

export enum Role {
  Admin = 'admin',
  User = 'user'
}

export type Round = {
  __typename?: 'Round';
  complete: Scalars['Boolean']['output'];
  create_time: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  number: Scalars['Int']['output'];
  score_caches: Array<ScoreCache>;
  statuses: Array<Status>;
  update_time: Scalars['Time']['output'];
};

export type Rubric = {
  __typename?: 'Rubric';
  fields: Array<RubricField>;
  notes?: Maybe<Scalars['String']['output']>;
};

export type RubricField = {
  __typename?: 'RubricField';
  name: Scalars['String']['output'];
  notes?: Maybe<Scalars['String']['output']>;
  score: Scalars['Int']['output'];
};

export type RubricFieldInput = {
  name: Scalars['String']['input'];
  notes?: InputMaybe<Scalars['String']['input']>;
  score: Scalars['Int']['input'];
};

export type RubricInput = {
  fields: Array<RubricFieldInput>;
  notes?: InputMaybe<Scalars['String']['input']>;
};

export type RubricTemplate = {
  __typename?: 'RubricTemplate';
  fields: Array<RubricTemplateField>;
  max_score: Scalars['Int']['output'];
};

export type RubricTemplateField = {
  __typename?: 'RubricTemplateField';
  max_score: Scalars['Int']['output'];
  name: Scalars['String']['output'];
};

export type RubricTemplateFieldInput = {
  max_score: Scalars['Int']['input'];
  name: Scalars['String']['input'];
};

export type RubricTemplateInput = {
  fields: Array<RubricTemplateFieldInput>;
  max_score: Scalars['Int']['input'];
};

export type Score = {
  __typename?: 'Score';
  score: Scalars['Int']['output'];
  user: User;
};

export type ScoreCache = {
  __typename?: 'ScoreCache';
  create_time: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  points: Scalars['Int']['output'];
  round: Round;
  round_id: Scalars['ID']['output'];
  update_time: Scalars['Time']['output'];
  user: User;
  user_id: Scalars['ID']['output'];
};

export type Scoreboard = {
  __typename?: 'Scoreboard';
  checks: Array<Check>;
  round: Round;
  scores: Array<Maybe<Score>>;
  statuses: Array<Array<Maybe<Status>>>;
  teams: Array<User>;
};

export type Source = {
  __typename?: 'Source';
  name: Scalars['String']['output'];
  schema: Scalars['String']['output'];
};

export type Status = {
  __typename?: 'Status';
  check: Check;
  check_id: Scalars['ID']['output'];
  create_time: Scalars['Time']['output'];
  error?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  points: Scalars['Int']['output'];
  round: Round;
  round_id: Scalars['ID']['output'];
  status: StatusEnum;
  update_time: Scalars['Time']['output'];
  user: User;
  user_id: Scalars['ID']['output'];
};

export enum StatusEnum {
  Down = 'down',
  Unknown = 'unknown',
  Up = 'up'
}

export type Subscription = {
  __typename?: 'Subscription';
  engineState: EngineState;
  globalNotification: Notification;
  latestRound: Round;
  scoreboardUpdate: Scoreboard;
};

export type User = {
  __typename?: 'User';
  configs: Array<Config>;
  create_time: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  inject_submissions: Array<InjectSubmission>;
  number?: Maybe<Scalars['Int']['output']>;
  role: Role;
  score_caches: Array<ScoreCache>;
  statuses: Array<Status>;
  update_time: Scalars['Time']['output'];
  username: Scalars['String']['output'];
};


export const MeDocument = gql`
    query Me {
  me {
    id
    username
    role
    number
  }
}
    `;

/**
 * __useMeQuery__
 *
 * To run a query within a React component, call `useMeQuery` and pass it any options that fit your needs.
 * When your component renders, `useMeQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useMeQuery({
 *   variables: {
 *   },
 * });
 */
export function useMeQuery(baseOptions?: Apollo.QueryHookOptions<MeQuery, MeQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<MeQuery, MeQueryVariables>(MeDocument, options);
      }
export function useMeLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<MeQuery, MeQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<MeQuery, MeQueryVariables>(MeDocument, options);
        }
export function useMeSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<MeQuery, MeQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<MeQuery, MeQueryVariables>(MeDocument, options);
        }
export type MeQueryHookResult = ReturnType<typeof useMeQuery>;
export type MeLazyQueryHookResult = ReturnType<typeof useMeLazyQuery>;
export type MeSuspenseQueryHookResult = ReturnType<typeof useMeSuspenseQuery>;
export type MeQueryResult = Apollo.QueryResult<MeQuery, MeQueryVariables>;
export const LoginDocument = gql`
    mutation Login($username: String!, $password: String!) {
  login(username: $username, password: $password) {
    name
    token
    expires
    path
    domain
    secure
    httpOnly
  }
}
    `;
export type LoginMutationFn = Apollo.MutationFunction<LoginMutation, LoginMutationVariables>;

/**
 * __useLoginMutation__
 *
 * To run a mutation, you first call `useLoginMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useLoginMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [loginMutation, { data, loading, error }] = useLoginMutation({
 *   variables: {
 *      username: // value for 'username'
 *      password: // value for 'password'
 *   },
 * });
 */
export function useLoginMutation(baseOptions?: Apollo.MutationHookOptions<LoginMutation, LoginMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<LoginMutation, LoginMutationVariables>(LoginDocument, options);
      }
export type LoginMutationHookResult = ReturnType<typeof useLoginMutation>;
export type LoginMutationResult = Apollo.MutationResult<LoginMutation>;
export type LoginMutationOptions = Apollo.BaseMutationOptions<LoginMutation, LoginMutationVariables>;
export const ChangePasswordDocument = gql`
    mutation ChangePassword($oldPassword: String!, $newPassword: String!) {
  changePassword(oldPassword: $oldPassword, newPassword: $newPassword)
}
    `;
export type ChangePasswordMutationFn = Apollo.MutationFunction<ChangePasswordMutation, ChangePasswordMutationVariables>;

/**
 * __useChangePasswordMutation__
 *
 * To run a mutation, you first call `useChangePasswordMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useChangePasswordMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [changePasswordMutation, { data, loading, error }] = useChangePasswordMutation({
 *   variables: {
 *      oldPassword: // value for 'oldPassword'
 *      newPassword: // value for 'newPassword'
 *   },
 * });
 */
export function useChangePasswordMutation(baseOptions?: Apollo.MutationHookOptions<ChangePasswordMutation, ChangePasswordMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<ChangePasswordMutation, ChangePasswordMutationVariables>(ChangePasswordDocument, options);
      }
export type ChangePasswordMutationHookResult = ReturnType<typeof useChangePasswordMutation>;
export type ChangePasswordMutationResult = Apollo.MutationResult<ChangePasswordMutation>;
export type ChangePasswordMutationOptions = Apollo.BaseMutationOptions<ChangePasswordMutation, ChangePasswordMutationVariables>;
export const ChecksDocument = gql`
    query Checks {
  checks {
    id
    name
    weight
    config
    editable_fields
    source {
      name
      schema
    }
  }
  sources {
    name
    schema
  }
}
    `;

/**
 * __useChecksQuery__
 *
 * To run a query within a React component, call `useChecksQuery` and pass it any options that fit your needs.
 * When your component renders, `useChecksQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useChecksQuery({
 *   variables: {
 *   },
 * });
 */
export function useChecksQuery(baseOptions?: Apollo.QueryHookOptions<ChecksQuery, ChecksQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ChecksQuery, ChecksQueryVariables>(ChecksDocument, options);
      }
export function useChecksLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ChecksQuery, ChecksQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ChecksQuery, ChecksQueryVariables>(ChecksDocument, options);
        }
export function useChecksSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<ChecksQuery, ChecksQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<ChecksQuery, ChecksQueryVariables>(ChecksDocument, options);
        }
export type ChecksQueryHookResult = ReturnType<typeof useChecksQuery>;
export type ChecksLazyQueryHookResult = ReturnType<typeof useChecksLazyQuery>;
export type ChecksSuspenseQueryHookResult = ReturnType<typeof useChecksSuspenseQuery>;
export type ChecksQueryResult = Apollo.QueryResult<ChecksQuery, ChecksQueryVariables>;
export const CreateCheckDocument = gql`
    mutation CreateCheck($name: String!, $weight: Int!, $source: String!, $config: JSON!, $editable_fields: [String!]!) {
  createCheck(
    name: $name
    weight: $weight
    source: $source
    config: $config
    editable_fields: $editable_fields
  ) {
    id
    name
    source {
      name
      schema
    }
  }
}
    `;
export type CreateCheckMutationFn = Apollo.MutationFunction<CreateCheckMutation, CreateCheckMutationVariables>;

/**
 * __useCreateCheckMutation__
 *
 * To run a mutation, you first call `useCreateCheckMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateCheckMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createCheckMutation, { data, loading, error }] = useCreateCheckMutation({
 *   variables: {
 *      name: // value for 'name'
 *      weight: // value for 'weight'
 *      source: // value for 'source'
 *      config: // value for 'config'
 *      editable_fields: // value for 'editable_fields'
 *   },
 * });
 */
export function useCreateCheckMutation(baseOptions?: Apollo.MutationHookOptions<CreateCheckMutation, CreateCheckMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateCheckMutation, CreateCheckMutationVariables>(CreateCheckDocument, options);
      }
export type CreateCheckMutationHookResult = ReturnType<typeof useCreateCheckMutation>;
export type CreateCheckMutationResult = Apollo.MutationResult<CreateCheckMutation>;
export type CreateCheckMutationOptions = Apollo.BaseMutationOptions<CreateCheckMutation, CreateCheckMutationVariables>;
export const UpdateCheckDocument = gql`
    mutation UpdateCheck($id: ID!, $name: String, $weight: Int, $config: JSON, $editable_fields: [String!]) {
  updateCheck(
    id: $id
    name: $name
    weight: $weight
    config: $config
    editable_fields: $editable_fields
  ) {
    id
    name
    source {
      name
      schema
    }
  }
}
    `;
export type UpdateCheckMutationFn = Apollo.MutationFunction<UpdateCheckMutation, UpdateCheckMutationVariables>;

/**
 * __useUpdateCheckMutation__
 *
 * To run a mutation, you first call `useUpdateCheckMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateCheckMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateCheckMutation, { data, loading, error }] = useUpdateCheckMutation({
 *   variables: {
 *      id: // value for 'id'
 *      name: // value for 'name'
 *      weight: // value for 'weight'
 *      config: // value for 'config'
 *      editable_fields: // value for 'editable_fields'
 *   },
 * });
 */
export function useUpdateCheckMutation(baseOptions?: Apollo.MutationHookOptions<UpdateCheckMutation, UpdateCheckMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateCheckMutation, UpdateCheckMutationVariables>(UpdateCheckDocument, options);
      }
export type UpdateCheckMutationHookResult = ReturnType<typeof useUpdateCheckMutation>;
export type UpdateCheckMutationResult = Apollo.MutationResult<UpdateCheckMutation>;
export type UpdateCheckMutationOptions = Apollo.BaseMutationOptions<UpdateCheckMutation, UpdateCheckMutationVariables>;
export const DeleteCheckDocument = gql`
    mutation DeleteCheck($id: ID!) {
  deleteCheck(id: $id)
}
    `;
export type DeleteCheckMutationFn = Apollo.MutationFunction<DeleteCheckMutation, DeleteCheckMutationVariables>;

/**
 * __useDeleteCheckMutation__
 *
 * To run a mutation, you first call `useDeleteCheckMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteCheckMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteCheckMutation, { data, loading, error }] = useDeleteCheckMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteCheckMutation(baseOptions?: Apollo.MutationHookOptions<DeleteCheckMutation, DeleteCheckMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteCheckMutation, DeleteCheckMutationVariables>(DeleteCheckDocument, options);
      }
export type DeleteCheckMutationHookResult = ReturnType<typeof useDeleteCheckMutation>;
export type DeleteCheckMutationResult = Apollo.MutationResult<DeleteCheckMutation>;
export type DeleteCheckMutationOptions = Apollo.BaseMutationOptions<DeleteCheckMutation, DeleteCheckMutationVariables>;
export const UsersDocument = gql`
    query Users {
  users {
    id
    username
    role
    number
  }
}
    `;

/**
 * __useUsersQuery__
 *
 * To run a query within a React component, call `useUsersQuery` and pass it any options that fit your needs.
 * When your component renders, `useUsersQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useUsersQuery({
 *   variables: {
 *   },
 * });
 */
export function useUsersQuery(baseOptions?: Apollo.QueryHookOptions<UsersQuery, UsersQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<UsersQuery, UsersQueryVariables>(UsersDocument, options);
      }
export function useUsersLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<UsersQuery, UsersQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<UsersQuery, UsersQueryVariables>(UsersDocument, options);
        }
export function useUsersSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<UsersQuery, UsersQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<UsersQuery, UsersQueryVariables>(UsersDocument, options);
        }
export type UsersQueryHookResult = ReturnType<typeof useUsersQuery>;
export type UsersLazyQueryHookResult = ReturnType<typeof useUsersLazyQuery>;
export type UsersSuspenseQueryHookResult = ReturnType<typeof useUsersSuspenseQuery>;
export type UsersQueryResult = Apollo.QueryResult<UsersQuery, UsersQueryVariables>;
export const CreateUserDocument = gql`
    mutation CreateUser($username: String!, $password: String!, $role: Role!, $number: Int) {
  createUser(
    username: $username
    password: $password
    role: $role
    number: $number
  ) {
    id
    username
    role
    number
  }
}
    `;
export type CreateUserMutationFn = Apollo.MutationFunction<CreateUserMutation, CreateUserMutationVariables>;

/**
 * __useCreateUserMutation__
 *
 * To run a mutation, you first call `useCreateUserMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateUserMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createUserMutation, { data, loading, error }] = useCreateUserMutation({
 *   variables: {
 *      username: // value for 'username'
 *      password: // value for 'password'
 *      role: // value for 'role'
 *      number: // value for 'number'
 *   },
 * });
 */
export function useCreateUserMutation(baseOptions?: Apollo.MutationHookOptions<CreateUserMutation, CreateUserMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateUserMutation, CreateUserMutationVariables>(CreateUserDocument, options);
      }
export type CreateUserMutationHookResult = ReturnType<typeof useCreateUserMutation>;
export type CreateUserMutationResult = Apollo.MutationResult<CreateUserMutation>;
export type CreateUserMutationOptions = Apollo.BaseMutationOptions<CreateUserMutation, CreateUserMutationVariables>;
export const UpdateUserDocument = gql`
    mutation UpdateUser($id: ID!, $username: String, $password: String, $number: Int) {
  updateUser(id: $id, username: $username, password: $password, number: $number) {
    id
    username
    number
  }
}
    `;
export type UpdateUserMutationFn = Apollo.MutationFunction<UpdateUserMutation, UpdateUserMutationVariables>;

/**
 * __useUpdateUserMutation__
 *
 * To run a mutation, you first call `useUpdateUserMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateUserMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateUserMutation, { data, loading, error }] = useUpdateUserMutation({
 *   variables: {
 *      id: // value for 'id'
 *      username: // value for 'username'
 *      password: // value for 'password'
 *      number: // value for 'number'
 *   },
 * });
 */
export function useUpdateUserMutation(baseOptions?: Apollo.MutationHookOptions<UpdateUserMutation, UpdateUserMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateUserMutation, UpdateUserMutationVariables>(UpdateUserDocument, options);
      }
export type UpdateUserMutationHookResult = ReturnType<typeof useUpdateUserMutation>;
export type UpdateUserMutationResult = Apollo.MutationResult<UpdateUserMutation>;
export type UpdateUserMutationOptions = Apollo.BaseMutationOptions<UpdateUserMutation, UpdateUserMutationVariables>;
export const DeleteUserDocument = gql`
    mutation DeleteUser($id: ID!) {
  deleteUser(id: $id)
}
    `;
export type DeleteUserMutationFn = Apollo.MutationFunction<DeleteUserMutation, DeleteUserMutationVariables>;

/**
 * __useDeleteUserMutation__
 *
 * To run a mutation, you first call `useDeleteUserMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteUserMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteUserMutation, { data, loading, error }] = useDeleteUserMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteUserMutation(baseOptions?: Apollo.MutationHookOptions<DeleteUserMutation, DeleteUserMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteUserMutation, DeleteUserMutationVariables>(DeleteUserDocument, options);
      }
export type DeleteUserMutationHookResult = ReturnType<typeof useDeleteUserMutation>;
export type DeleteUserMutationResult = Apollo.MutationResult<DeleteUserMutation>;
export type DeleteUserMutationOptions = Apollo.BaseMutationOptions<DeleteUserMutation, DeleteUserMutationVariables>;
export const GlobalNotificationDocument = gql`
    subscription GlobalNotification {
  globalNotification {
    message
    type
  }
}
    `;

/**
 * __useGlobalNotificationSubscription__
 *
 * To run a query within a React component, call `useGlobalNotificationSubscription` and pass it any options that fit your needs.
 * When your component renders, `useGlobalNotificationSubscription` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGlobalNotificationSubscription({
 *   variables: {
 *   },
 * });
 */
export function useGlobalNotificationSubscription(baseOptions?: Apollo.SubscriptionHookOptions<GlobalNotificationSubscription, GlobalNotificationSubscriptionVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useSubscription<GlobalNotificationSubscription, GlobalNotificationSubscriptionVariables>(GlobalNotificationDocument, options);
      }
export type GlobalNotificationSubscriptionHookResult = ReturnType<typeof useGlobalNotificationSubscription>;
export type GlobalNotificationSubscriptionResult = Apollo.SubscriptionResult<GlobalNotificationSubscription>;
export const EngineStateDocument = gql`
    subscription EngineState {
  engineState
}
    `;

/**
 * __useEngineStateSubscription__
 *
 * To run a query within a React component, call `useEngineStateSubscription` and pass it any options that fit your needs.
 * When your component renders, `useEngineStateSubscription` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useEngineStateSubscription({
 *   variables: {
 *   },
 * });
 */
export function useEngineStateSubscription(baseOptions?: Apollo.SubscriptionHookOptions<EngineStateSubscription, EngineStateSubscriptionVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useSubscription<EngineStateSubscription, EngineStateSubscriptionVariables>(EngineStateDocument, options);
      }
export type EngineStateSubscriptionHookResult = ReturnType<typeof useEngineStateSubscription>;
export type EngineStateSubscriptionResult = Apollo.SubscriptionResult<EngineStateSubscription>;
export const StartEngineDocument = gql`
    mutation StartEngine {
  startEngine
}
    `;
export type StartEngineMutationFn = Apollo.MutationFunction<StartEngineMutation, StartEngineMutationVariables>;

/**
 * __useStartEngineMutation__
 *
 * To run a mutation, you first call `useStartEngineMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useStartEngineMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [startEngineMutation, { data, loading, error }] = useStartEngineMutation({
 *   variables: {
 *   },
 * });
 */
export function useStartEngineMutation(baseOptions?: Apollo.MutationHookOptions<StartEngineMutation, StartEngineMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<StartEngineMutation, StartEngineMutationVariables>(StartEngineDocument, options);
      }
export type StartEngineMutationHookResult = ReturnType<typeof useStartEngineMutation>;
export type StartEngineMutationResult = Apollo.MutationResult<StartEngineMutation>;
export type StartEngineMutationOptions = Apollo.BaseMutationOptions<StartEngineMutation, StartEngineMutationVariables>;
export const StopEngineDocument = gql`
    mutation StopEngine {
  stopEngine
}
    `;
export type StopEngineMutationFn = Apollo.MutationFunction<StopEngineMutation, StopEngineMutationVariables>;

/**
 * __useStopEngineMutation__
 *
 * To run a mutation, you first call `useStopEngineMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useStopEngineMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [stopEngineMutation, { data, loading, error }] = useStopEngineMutation({
 *   variables: {
 *   },
 * });
 */
export function useStopEngineMutation(baseOptions?: Apollo.MutationHookOptions<StopEngineMutation, StopEngineMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<StopEngineMutation, StopEngineMutationVariables>(StopEngineDocument, options);
      }
export type StopEngineMutationHookResult = ReturnType<typeof useStopEngineMutation>;
export type StopEngineMutationResult = Apollo.MutationResult<StopEngineMutation>;
export type StopEngineMutationOptions = Apollo.BaseMutationOptions<StopEngineMutation, StopEngineMutationVariables>;
export const SendGlobalNotificationDocument = gql`
    mutation SendGlobalNotification($message: String!, $type: NotificationType!) {
  sendGlobalNotification(message: $message, type: $type)
}
    `;
export type SendGlobalNotificationMutationFn = Apollo.MutationFunction<SendGlobalNotificationMutation, SendGlobalNotificationMutationVariables>;

/**
 * __useSendGlobalNotificationMutation__
 *
 * To run a mutation, you first call `useSendGlobalNotificationMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useSendGlobalNotificationMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [sendGlobalNotificationMutation, { data, loading, error }] = useSendGlobalNotificationMutation({
 *   variables: {
 *      message: // value for 'message'
 *      type: // value for 'type'
 *   },
 * });
 */
export function useSendGlobalNotificationMutation(baseOptions?: Apollo.MutationHookOptions<SendGlobalNotificationMutation, SendGlobalNotificationMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<SendGlobalNotificationMutation, SendGlobalNotificationMutationVariables>(SendGlobalNotificationDocument, options);
      }
export type SendGlobalNotificationMutationHookResult = ReturnType<typeof useSendGlobalNotificationMutation>;
export type SendGlobalNotificationMutationResult = Apollo.MutationResult<SendGlobalNotificationMutation>;
export type SendGlobalNotificationMutationOptions = Apollo.BaseMutationOptions<SendGlobalNotificationMutation, SendGlobalNotificationMutationVariables>;
export const AdminLoginDocument = gql`
    mutation AdminLogin($id: ID!) {
  adminLogin(id: $id) {
    name
    token
    expires
    path
    domain
    secure
    httpOnly
  }
}
    `;
export type AdminLoginMutationFn = Apollo.MutationFunction<AdminLoginMutation, AdminLoginMutationVariables>;

/**
 * __useAdminLoginMutation__
 *
 * To run a mutation, you first call `useAdminLoginMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAdminLoginMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [adminLoginMutation, { data, loading, error }] = useAdminLoginMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useAdminLoginMutation(baseOptions?: Apollo.MutationHookOptions<AdminLoginMutation, AdminLoginMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AdminLoginMutation, AdminLoginMutationVariables>(AdminLoginDocument, options);
      }
export type AdminLoginMutationHookResult = ReturnType<typeof useAdminLoginMutation>;
export type AdminLoginMutationResult = Apollo.MutationResult<AdminLoginMutation>;
export type AdminLoginMutationOptions = Apollo.BaseMutationOptions<AdminLoginMutation, AdminLoginMutationVariables>;
export const AdminBecomeDocument = gql`
    mutation AdminBecome($id: ID!) {
  adminBecome(id: $id) {
    name
    token
    expires
    path
    domain
    secure
    httpOnly
  }
}
    `;
export type AdminBecomeMutationFn = Apollo.MutationFunction<AdminBecomeMutation, AdminBecomeMutationVariables>;

/**
 * __useAdminBecomeMutation__
 *
 * To run a mutation, you first call `useAdminBecomeMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAdminBecomeMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [adminBecomeMutation, { data, loading, error }] = useAdminBecomeMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useAdminBecomeMutation(baseOptions?: Apollo.MutationHookOptions<AdminBecomeMutation, AdminBecomeMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AdminBecomeMutation, AdminBecomeMutationVariables>(AdminBecomeDocument, options);
      }
export type AdminBecomeMutationHookResult = ReturnType<typeof useAdminBecomeMutation>;
export type AdminBecomeMutationResult = Apollo.MutationResult<AdminBecomeMutation>;
export type AdminBecomeMutationOptions = Apollo.BaseMutationOptions<AdminBecomeMutation, AdminBecomeMutationVariables>;
export const ConfigsDocument = gql`
    query Configs {
  configs {
    id
    check {
      name
      weight
      source {
        name
        schema
      }
    }
    config
  }
}
    `;

/**
 * __useConfigsQuery__
 *
 * To run a query within a React component, call `useConfigsQuery` and pass it any options that fit your needs.
 * When your component renders, `useConfigsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useConfigsQuery({
 *   variables: {
 *   },
 * });
 */
export function useConfigsQuery(baseOptions?: Apollo.QueryHookOptions<ConfigsQuery, ConfigsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ConfigsQuery, ConfigsQueryVariables>(ConfigsDocument, options);
      }
export function useConfigsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ConfigsQuery, ConfigsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ConfigsQuery, ConfigsQueryVariables>(ConfigsDocument, options);
        }
export function useConfigsSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<ConfigsQuery, ConfigsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<ConfigsQuery, ConfigsQueryVariables>(ConfigsDocument, options);
        }
export type ConfigsQueryHookResult = ReturnType<typeof useConfigsQuery>;
export type ConfigsLazyQueryHookResult = ReturnType<typeof useConfigsLazyQuery>;
export type ConfigsSuspenseQueryHookResult = ReturnType<typeof useConfigsSuspenseQuery>;
export type ConfigsQueryResult = Apollo.QueryResult<ConfigsQuery, ConfigsQueryVariables>;
export const EditConfigDocument = gql`
    mutation EditConfig($id: ID!, $config: JSON!) {
  editConfig(id: $id, config: $config) {
    id
  }
}
    `;
export type EditConfigMutationFn = Apollo.MutationFunction<EditConfigMutation, EditConfigMutationVariables>;

/**
 * __useEditConfigMutation__
 *
 * To run a mutation, you first call `useEditConfigMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useEditConfigMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [editConfigMutation, { data, loading, error }] = useEditConfigMutation({
 *   variables: {
 *      id: // value for 'id'
 *      config: // value for 'config'
 *   },
 * });
 */
export function useEditConfigMutation(baseOptions?: Apollo.MutationHookOptions<EditConfigMutation, EditConfigMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<EditConfigMutation, EditConfigMutationVariables>(EditConfigDocument, options);
      }
export type EditConfigMutationHookResult = ReturnType<typeof useEditConfigMutation>;
export type EditConfigMutationResult = Apollo.MutationResult<EditConfigMutation>;
export type EditConfigMutationOptions = Apollo.BaseMutationOptions<EditConfigMutation, EditConfigMutationVariables>;
export const ScoreboardDocument = gql`
    query Scoreboard($round: Int) {
  scoreboard(round: $round) {
    round {
      number
    }
    teams {
      username
      number
    }
    checks {
      name
    }
    statuses {
      error
      status
      update_time
    }
    scores {
      user {
        username
        number
      }
      score
    }
  }
}
    `;

/**
 * __useScoreboardQuery__
 *
 * To run a query within a React component, call `useScoreboardQuery` and pass it any options that fit your needs.
 * When your component renders, `useScoreboardQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useScoreboardQuery({
 *   variables: {
 *      round: // value for 'round'
 *   },
 * });
 */
export function useScoreboardQuery(baseOptions?: Apollo.QueryHookOptions<ScoreboardQuery, ScoreboardQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ScoreboardQuery, ScoreboardQueryVariables>(ScoreboardDocument, options);
      }
export function useScoreboardLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ScoreboardQuery, ScoreboardQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ScoreboardQuery, ScoreboardQueryVariables>(ScoreboardDocument, options);
        }
export function useScoreboardSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<ScoreboardQuery, ScoreboardQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<ScoreboardQuery, ScoreboardQueryVariables>(ScoreboardDocument, options);
        }
export type ScoreboardQueryHookResult = ReturnType<typeof useScoreboardQuery>;
export type ScoreboardLazyQueryHookResult = ReturnType<typeof useScoreboardLazyQuery>;
export type ScoreboardSuspenseQueryHookResult = ReturnType<typeof useScoreboardSuspenseQuery>;
export type ScoreboardQueryResult = Apollo.QueryResult<ScoreboardQuery, ScoreboardQueryVariables>;
export const ScoreboardUpdateDocument = gql`
    subscription ScoreboardUpdate {
  scoreboardUpdate {
    round {
      number
    }
    teams {
      username
      number
    }
    checks {
      name
    }
    statuses {
      error
      status
      update_time
    }
    scores {
      user {
        username
        number
      }
      score
    }
  }
}
    `;

/**
 * __useScoreboardUpdateSubscription__
 *
 * To run a query within a React component, call `useScoreboardUpdateSubscription` and pass it any options that fit your needs.
 * When your component renders, `useScoreboardUpdateSubscription` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useScoreboardUpdateSubscription({
 *   variables: {
 *   },
 * });
 */
export function useScoreboardUpdateSubscription(baseOptions?: Apollo.SubscriptionHookOptions<ScoreboardUpdateSubscription, ScoreboardUpdateSubscriptionVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useSubscription<ScoreboardUpdateSubscription, ScoreboardUpdateSubscriptionVariables>(ScoreboardUpdateDocument, options);
      }
export type ScoreboardUpdateSubscriptionHookResult = ReturnType<typeof useScoreboardUpdateSubscription>;
export type ScoreboardUpdateSubscriptionResult = Apollo.SubscriptionResult<ScoreboardUpdateSubscription>;
export const LatestRoundDocument = gql`
    subscription LatestRound {
  latestRound {
    number
  }
}
    `;

/**
 * __useLatestRoundSubscription__
 *
 * To run a query within a React component, call `useLatestRoundSubscription` and pass it any options that fit your needs.
 * When your component renders, `useLatestRoundSubscription` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useLatestRoundSubscription({
 *   variables: {
 *   },
 * });
 */
export function useLatestRoundSubscription(baseOptions?: Apollo.SubscriptionHookOptions<LatestRoundSubscription, LatestRoundSubscriptionVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useSubscription<LatestRoundSubscription, LatestRoundSubscriptionVariables>(LatestRoundDocument, options);
      }
export type LatestRoundSubscriptionHookResult = ReturnType<typeof useLatestRoundSubscription>;
export type LatestRoundSubscriptionResult = Apollo.SubscriptionResult<LatestRoundSubscription>;
export const CreateInjectDocument = gql`
    mutation CreateInject($title: String!, $start_time: Time!, $end_time: Time!, $files: [Upload!]!, $rubric: RubricTemplateInput!) {
  createInject(
    title: $title
    start_time: $start_time
    end_time: $end_time
    files: $files
    rubric: $rubric
  ) {
    id
  }
}
    `;
export type CreateInjectMutationFn = Apollo.MutationFunction<CreateInjectMutation, CreateInjectMutationVariables>;

/**
 * __useCreateInjectMutation__
 *
 * To run a mutation, you first call `useCreateInjectMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateInjectMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createInjectMutation, { data, loading, error }] = useCreateInjectMutation({
 *   variables: {
 *      title: // value for 'title'
 *      start_time: // value for 'start_time'
 *      end_time: // value for 'end_time'
 *      files: // value for 'files'
 *      rubric: // value for 'rubric'
 *   },
 * });
 */
export function useCreateInjectMutation(baseOptions?: Apollo.MutationHookOptions<CreateInjectMutation, CreateInjectMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateInjectMutation, CreateInjectMutationVariables>(CreateInjectDocument, options);
      }
export type CreateInjectMutationHookResult = ReturnType<typeof useCreateInjectMutation>;
export type CreateInjectMutationResult = Apollo.MutationResult<CreateInjectMutation>;
export type CreateInjectMutationOptions = Apollo.BaseMutationOptions<CreateInjectMutation, CreateInjectMutationVariables>;
export const InjectsDocument = gql`
    query Injects {
  injects {
    id
    title
    start_time
    end_time
    files {
      id
      name
      url
    }
    rubric {
      max_score
      fields {
        name
        max_score
      }
    }
    submissions {
      id
      create_time
      update_time
      files {
        id
        name
        url
      }
      rubric {
        fields {
          name
          score
          notes
        }
        notes
      }
      notes
      graded
    }
  }
}
    `;

/**
 * __useInjectsQuery__
 *
 * To run a query within a React component, call `useInjectsQuery` and pass it any options that fit your needs.
 * When your component renders, `useInjectsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useInjectsQuery({
 *   variables: {
 *   },
 * });
 */
export function useInjectsQuery(baseOptions?: Apollo.QueryHookOptions<InjectsQuery, InjectsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<InjectsQuery, InjectsQueryVariables>(InjectsDocument, options);
      }
export function useInjectsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<InjectsQuery, InjectsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<InjectsQuery, InjectsQueryVariables>(InjectsDocument, options);
        }
export function useInjectsSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<InjectsQuery, InjectsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<InjectsQuery, InjectsQueryVariables>(InjectsDocument, options);
        }
export type InjectsQueryHookResult = ReturnType<typeof useInjectsQuery>;
export type InjectsLazyQueryHookResult = ReturnType<typeof useInjectsLazyQuery>;
export type InjectsSuspenseQueryHookResult = ReturnType<typeof useInjectsSuspenseQuery>;
export type InjectsQueryResult = Apollo.QueryResult<InjectsQuery, InjectsQueryVariables>;
export const UpdateInjectDocument = gql`
    mutation UpdateInject($id: ID!, $title: String, $start_time: Time, $end_time: Time, $delete_files: [ID!], $add_files: [Upload!], $rubric: RubricTemplateInput) {
  updateInject(
    id: $id
    title: $title
    start_time: $start_time
    end_time: $end_time
    delete_files: $delete_files
    add_files: $add_files
    rubric: $rubric
  ) {
    id
  }
}
    `;
export type UpdateInjectMutationFn = Apollo.MutationFunction<UpdateInjectMutation, UpdateInjectMutationVariables>;

/**
 * __useUpdateInjectMutation__
 *
 * To run a mutation, you first call `useUpdateInjectMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateInjectMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateInjectMutation, { data, loading, error }] = useUpdateInjectMutation({
 *   variables: {
 *      id: // value for 'id'
 *      title: // value for 'title'
 *      start_time: // value for 'start_time'
 *      end_time: // value for 'end_time'
 *      delete_files: // value for 'delete_files'
 *      add_files: // value for 'add_files'
 *      rubric: // value for 'rubric'
 *   },
 * });
 */
export function useUpdateInjectMutation(baseOptions?: Apollo.MutationHookOptions<UpdateInjectMutation, UpdateInjectMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateInjectMutation, UpdateInjectMutationVariables>(UpdateInjectDocument, options);
      }
export type UpdateInjectMutationHookResult = ReturnType<typeof useUpdateInjectMutation>;
export type UpdateInjectMutationResult = Apollo.MutationResult<UpdateInjectMutation>;
export type UpdateInjectMutationOptions = Apollo.BaseMutationOptions<UpdateInjectMutation, UpdateInjectMutationVariables>;
export const DeleteInjectDocument = gql`
    mutation DeleteInject($id: ID!) {
  deleteInject(id: $id)
}
    `;
export type DeleteInjectMutationFn = Apollo.MutationFunction<DeleteInjectMutation, DeleteInjectMutationVariables>;

/**
 * __useDeleteInjectMutation__
 *
 * To run a mutation, you first call `useDeleteInjectMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteInjectMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteInjectMutation, { data, loading, error }] = useDeleteInjectMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteInjectMutation(baseOptions?: Apollo.MutationHookOptions<DeleteInjectMutation, DeleteInjectMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteInjectMutation, DeleteInjectMutationVariables>(DeleteInjectDocument, options);
      }
export type DeleteInjectMutationHookResult = ReturnType<typeof useDeleteInjectMutation>;
export type DeleteInjectMutationResult = Apollo.MutationResult<DeleteInjectMutation>;
export type DeleteInjectMutationOptions = Apollo.BaseMutationOptions<DeleteInjectMutation, DeleteInjectMutationVariables>;
export const SubmitInjectDocument = gql`
    mutation SubmitInject($id: ID!, $files: [Upload!]!, $notes: String!) {
  submitInject(injectID: $id, files: $files, notes: $notes) {
    id
  }
}
    `;
export type SubmitInjectMutationFn = Apollo.MutationFunction<SubmitInjectMutation, SubmitInjectMutationVariables>;

/**
 * __useSubmitInjectMutation__
 *
 * To run a mutation, you first call `useSubmitInjectMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useSubmitInjectMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [submitInjectMutation, { data, loading, error }] = useSubmitInjectMutation({
 *   variables: {
 *      id: // value for 'id'
 *      files: // value for 'files'
 *      notes: // value for 'notes'
 *   },
 * });
 */
export function useSubmitInjectMutation(baseOptions?: Apollo.MutationHookOptions<SubmitInjectMutation, SubmitInjectMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<SubmitInjectMutation, SubmitInjectMutationVariables>(SubmitInjectDocument, options);
      }
export type SubmitInjectMutationHookResult = ReturnType<typeof useSubmitInjectMutation>;
export type SubmitInjectMutationResult = Apollo.MutationResult<SubmitInjectMutation>;
export type SubmitInjectMutationOptions = Apollo.BaseMutationOptions<SubmitInjectMutation, SubmitInjectMutationVariables>;
export const SubmissionsDocument = gql`
    query Submissions($inject_id: ID!) {
  injectSubmissionsByUser(id: $inject_id) {
    user {
      username
      number
    }
    submissions {
      id
      create_time
      update_time
      files {
        id
        name
        url
      }
      user {
        id
        username
      }
      inject {
        id
        title
        start_time
        end_time
        create_time
        update_time
        rubric {
          fields {
            name
            max_score
          }
          max_score
        }
      }
      rubric {
        fields {
          name
          score
          notes
        }
        notes
      }
      graded
      notes
    }
  }
}
    `;

/**
 * __useSubmissionsQuery__
 *
 * To run a query within a React component, call `useSubmissionsQuery` and pass it any options that fit your needs.
 * When your component renders, `useSubmissionsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useSubmissionsQuery({
 *   variables: {
 *      inject_id: // value for 'inject_id'
 *   },
 * });
 */
export function useSubmissionsQuery(baseOptions: Apollo.QueryHookOptions<SubmissionsQuery, SubmissionsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<SubmissionsQuery, SubmissionsQueryVariables>(SubmissionsDocument, options);
      }
export function useSubmissionsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<SubmissionsQuery, SubmissionsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<SubmissionsQuery, SubmissionsQueryVariables>(SubmissionsDocument, options);
        }
export function useSubmissionsSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<SubmissionsQuery, SubmissionsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<SubmissionsQuery, SubmissionsQueryVariables>(SubmissionsDocument, options);
        }
export type SubmissionsQueryHookResult = ReturnType<typeof useSubmissionsQuery>;
export type SubmissionsLazyQueryHookResult = ReturnType<typeof useSubmissionsLazyQuery>;
export type SubmissionsSuspenseQueryHookResult = ReturnType<typeof useSubmissionsSuspenseQuery>;
export type SubmissionsQueryResult = Apollo.QueryResult<SubmissionsQuery, SubmissionsQueryVariables>;
export const GradeSubmissionDocument = gql`
    mutation GradeSubmission($submission_id: ID!, $rubric: RubricInput!) {
  gradeSubmission(submissionID: $submission_id, rubric: $rubric) {
    id
  }
}
    `;
export type GradeSubmissionMutationFn = Apollo.MutationFunction<GradeSubmissionMutation, GradeSubmissionMutationVariables>;

/**
 * __useGradeSubmissionMutation__
 *
 * To run a mutation, you first call `useGradeSubmissionMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useGradeSubmissionMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [gradeSubmissionMutation, { data, loading, error }] = useGradeSubmissionMutation({
 *   variables: {
 *      submission_id: // value for 'submission_id'
 *      rubric: // value for 'rubric'
 *   },
 * });
 */
export function useGradeSubmissionMutation(baseOptions?: Apollo.MutationHookOptions<GradeSubmissionMutation, GradeSubmissionMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<GradeSubmissionMutation, GradeSubmissionMutationVariables>(GradeSubmissionDocument, options);
      }
export type GradeSubmissionMutationHookResult = ReturnType<typeof useGradeSubmissionMutation>;
export type GradeSubmissionMutationResult = Apollo.MutationResult<GradeSubmissionMutation>;
export type GradeSubmissionMutationOptions = Apollo.BaseMutationOptions<GradeSubmissionMutation, GradeSubmissionMutationVariables>;
export type MeQueryVariables = Exact<{ [key: string]: never; }>;


export type MeQuery = { __typename?: 'Query', me?: { __typename?: 'User', id: string, username: string, role: Role, number?: number | null } | null };

export type LoginMutationVariables = Exact<{
  username: Scalars['String']['input'];
  password: Scalars['String']['input'];
}>;


export type LoginMutation = { __typename?: 'Mutation', login: { __typename?: 'LoginOutput', name: string, token: string, expires: number, path: string, domain: string, secure: boolean, httpOnly: boolean } };

export type ChangePasswordMutationVariables = Exact<{
  oldPassword: Scalars['String']['input'];
  newPassword: Scalars['String']['input'];
}>;


export type ChangePasswordMutation = { __typename?: 'Mutation', changePassword: boolean };

export type ChecksQueryVariables = Exact<{ [key: string]: never; }>;


export type ChecksQuery = { __typename?: 'Query', checks: Array<{ __typename?: 'Check', id: string, name: string, weight: number, config: any, editable_fields: Array<string>, source: { __typename?: 'Source', name: string, schema: string } }>, sources: Array<{ __typename?: 'Source', name: string, schema: string }> };

export type CreateCheckMutationVariables = Exact<{
  name: Scalars['String']['input'];
  weight: Scalars['Int']['input'];
  source: Scalars['String']['input'];
  config: Scalars['JSON']['input'];
  editable_fields: Array<Scalars['String']['input']> | Scalars['String']['input'];
}>;


export type CreateCheckMutation = { __typename?: 'Mutation', createCheck: { __typename?: 'Check', id: string, name: string, source: { __typename?: 'Source', name: string, schema: string } } };

export type UpdateCheckMutationVariables = Exact<{
  id: Scalars['ID']['input'];
  name?: InputMaybe<Scalars['String']['input']>;
  weight?: InputMaybe<Scalars['Int']['input']>;
  config?: InputMaybe<Scalars['JSON']['input']>;
  editable_fields?: InputMaybe<Array<Scalars['String']['input']> | Scalars['String']['input']>;
}>;


export type UpdateCheckMutation = { __typename?: 'Mutation', updateCheck: { __typename?: 'Check', id: string, name: string, source: { __typename?: 'Source', name: string, schema: string } } };

export type DeleteCheckMutationVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type DeleteCheckMutation = { __typename?: 'Mutation', deleteCheck: boolean };

export type UsersQueryVariables = Exact<{ [key: string]: never; }>;


export type UsersQuery = { __typename?: 'Query', users: Array<{ __typename?: 'User', id: string, username: string, role: Role, number?: number | null }> };

export type CreateUserMutationVariables = Exact<{
  username: Scalars['String']['input'];
  password: Scalars['String']['input'];
  role: Role;
  number?: InputMaybe<Scalars['Int']['input']>;
}>;


export type CreateUserMutation = { __typename?: 'Mutation', createUser: { __typename?: 'User', id: string, username: string, role: Role, number?: number | null } };

export type UpdateUserMutationVariables = Exact<{
  id: Scalars['ID']['input'];
  username?: InputMaybe<Scalars['String']['input']>;
  password?: InputMaybe<Scalars['String']['input']>;
  number?: InputMaybe<Scalars['Int']['input']>;
}>;


export type UpdateUserMutation = { __typename?: 'Mutation', updateUser: { __typename?: 'User', id: string, username: string, number?: number | null } };

export type DeleteUserMutationVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type DeleteUserMutation = { __typename?: 'Mutation', deleteUser: boolean };

export type GlobalNotificationSubscriptionVariables = Exact<{ [key: string]: never; }>;


export type GlobalNotificationSubscription = { __typename?: 'Subscription', globalNotification: { __typename?: 'Notification', message: string, type: NotificationType } };

export type EngineStateSubscriptionVariables = Exact<{ [key: string]: never; }>;


export type EngineStateSubscription = { __typename?: 'Subscription', engineState: EngineState };

export type StartEngineMutationVariables = Exact<{ [key: string]: never; }>;


export type StartEngineMutation = { __typename?: 'Mutation', startEngine: boolean };

export type StopEngineMutationVariables = Exact<{ [key: string]: never; }>;


export type StopEngineMutation = { __typename?: 'Mutation', stopEngine: boolean };

export type SendGlobalNotificationMutationVariables = Exact<{
  message: Scalars['String']['input'];
  type: NotificationType;
}>;


export type SendGlobalNotificationMutation = { __typename?: 'Mutation', sendGlobalNotification: boolean };

export type AdminLoginMutationVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type AdminLoginMutation = { __typename?: 'Mutation', adminLogin: { __typename?: 'LoginOutput', name: string, token: string, expires: number, path: string, domain: string, secure: boolean, httpOnly: boolean } };

export type AdminBecomeMutationVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type AdminBecomeMutation = { __typename?: 'Mutation', adminBecome: { __typename?: 'LoginOutput', name: string, token: string, expires: number, path: string, domain: string, secure: boolean, httpOnly: boolean } };

export type ConfigsQueryVariables = Exact<{ [key: string]: never; }>;


export type ConfigsQuery = { __typename?: 'Query', configs: Array<{ __typename?: 'Config', id: string, config: any, check: { __typename?: 'Check', name: string, weight: number, source: { __typename?: 'Source', name: string, schema: string } } }> };

export type EditConfigMutationVariables = Exact<{
  id: Scalars['ID']['input'];
  config: Scalars['JSON']['input'];
}>;


export type EditConfigMutation = { __typename?: 'Mutation', editConfig: { __typename?: 'Config', id: string } };

export type ScoreboardQueryVariables = Exact<{
  round?: InputMaybe<Scalars['Int']['input']>;
}>;


export type ScoreboardQuery = { __typename?: 'Query', scoreboard: { __typename?: 'Scoreboard', round: { __typename?: 'Round', number: number }, teams: Array<{ __typename?: 'User', username: string, number?: number | null }>, checks: Array<{ __typename?: 'Check', name: string }>, statuses: Array<Array<{ __typename?: 'Status', error?: string | null, status: StatusEnum, update_time: any } | null>>, scores: Array<{ __typename?: 'Score', score: number, user: { __typename?: 'User', username: string, number?: number | null } } | null> } };

export type ScoreboardUpdateSubscriptionVariables = Exact<{ [key: string]: never; }>;


export type ScoreboardUpdateSubscription = { __typename?: 'Subscription', scoreboardUpdate: { __typename?: 'Scoreboard', round: { __typename?: 'Round', number: number }, teams: Array<{ __typename?: 'User', username: string, number?: number | null }>, checks: Array<{ __typename?: 'Check', name: string }>, statuses: Array<Array<{ __typename?: 'Status', error?: string | null, status: StatusEnum, update_time: any } | null>>, scores: Array<{ __typename?: 'Score', score: number, user: { __typename?: 'User', username: string, number?: number | null } } | null> } };

export type LatestRoundSubscriptionVariables = Exact<{ [key: string]: never; }>;


export type LatestRoundSubscription = { __typename?: 'Subscription', latestRound: { __typename?: 'Round', number: number } };

export type CreateInjectMutationVariables = Exact<{
  title: Scalars['String']['input'];
  start_time: Scalars['Time']['input'];
  end_time: Scalars['Time']['input'];
  files: Array<Scalars['Upload']['input']> | Scalars['Upload']['input'];
  rubric: RubricTemplateInput;
}>;


export type CreateInjectMutation = { __typename?: 'Mutation', createInject: { __typename?: 'Inject', id: string } };

export type InjectsQueryVariables = Exact<{ [key: string]: never; }>;


export type InjectsQuery = { __typename?: 'Query', injects: Array<{ __typename?: 'Inject', id: string, title: string, start_time: any, end_time: any, files: Array<{ __typename?: 'File', id: string, name: string, url: string }>, rubric: { __typename?: 'RubricTemplate', max_score: number, fields: Array<{ __typename?: 'RubricTemplateField', name: string, max_score: number }> }, submissions: Array<{ __typename?: 'InjectSubmission', id: string, create_time: any, update_time: any, notes: string, graded: boolean, files: Array<{ __typename?: 'File', id: string, name: string, url: string }>, rubric?: { __typename?: 'Rubric', notes?: string | null, fields: Array<{ __typename?: 'RubricField', name: string, score: number, notes?: string | null }> } | null }> }> };

export type UpdateInjectMutationVariables = Exact<{
  id: Scalars['ID']['input'];
  title?: InputMaybe<Scalars['String']['input']>;
  start_time?: InputMaybe<Scalars['Time']['input']>;
  end_time?: InputMaybe<Scalars['Time']['input']>;
  delete_files?: InputMaybe<Array<Scalars['ID']['input']> | Scalars['ID']['input']>;
  add_files?: InputMaybe<Array<Scalars['Upload']['input']> | Scalars['Upload']['input']>;
  rubric?: InputMaybe<RubricTemplateInput>;
}>;


export type UpdateInjectMutation = { __typename?: 'Mutation', updateInject: { __typename?: 'Inject', id: string } };

export type DeleteInjectMutationVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type DeleteInjectMutation = { __typename?: 'Mutation', deleteInject: boolean };

export type SubmitInjectMutationVariables = Exact<{
  id: Scalars['ID']['input'];
  files: Array<Scalars['Upload']['input']> | Scalars['Upload']['input'];
  notes: Scalars['String']['input'];
}>;


export type SubmitInjectMutation = { __typename?: 'Mutation', submitInject: { __typename?: 'InjectSubmission', id: string } };

export type SubmissionsQueryVariables = Exact<{
  inject_id: Scalars['ID']['input'];
}>;


export type SubmissionsQuery = { __typename?: 'Query', injectSubmissionsByUser: Array<{ __typename?: 'InjectSubmissionByUser', user: { __typename?: 'User', username: string, number?: number | null }, submissions: Array<{ __typename?: 'InjectSubmission', id: string, create_time: any, update_time: any, graded: boolean, notes: string, files: Array<{ __typename?: 'File', id: string, name: string, url: string }>, user: { __typename?: 'User', id: string, username: string }, inject: { __typename?: 'Inject', id: string, title: string, start_time: any, end_time: any, create_time: any, update_time: any, rubric: { __typename?: 'RubricTemplate', max_score: number, fields: Array<{ __typename?: 'RubricTemplateField', name: string, max_score: number }> } }, rubric?: { __typename?: 'Rubric', notes?: string | null, fields: Array<{ __typename?: 'RubricField', name: string, score: number, notes?: string | null }> } | null }> }> };

export type GradeSubmissionMutationVariables = Exact<{
  submission_id: Scalars['ID']['input'];
  rubric: RubricInput;
}>;


export type GradeSubmissionMutation = { __typename?: 'Mutation', gradeSubmission: { __typename?: 'InjectSubmission', id: string } };
