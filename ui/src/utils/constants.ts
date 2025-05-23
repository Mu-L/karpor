export const SEVERITY_MAP = {
  Critical: {
    color: 'purple',
    text: 'Critical',
  },
  High: {
    color: 'error',
    text: 'High',
  },
  Medium: {
    color: 'warning',
    text: 'Medium',
  },
  Low: {
    color: 'success',
    text: 'Low',
  },
  Safe: {
    color: 'success',
    text: 'Safe',
  },
}

export const searchSqlPrefix = 'select * from resources'

export const whereKeywords = [
  'cluster',
  'kind',
  'namespace',
  'name',
  'content',
  '`labels.${key}`',
  '`annotations.${key}`',
]
export const operatorKeywords = ['=', 'like', 'contains']
export const kindCompletions = [
  `'Pod'`,
  `'Service'`,
  `'ReplicaSet'`,
  `'Deployment'`,
  `'StatefulSet'`,
  `'DaemonSet'`,
  `'Job'`,
  `'CronJob'`,
  `'Namespace'`,
  `'ServiceAccount'`,
  `'ConfigMap'`,
  `'Secret'`,
  `'PersistentVolume'`,
  `'PersistentVolumeClaim'`,
  `'Ingress'`,
  `'StorageClass'`,
  `'NetworkPolicy'`,
  `'ResourceQuota'`,
  `'LimitRange'`,
  `'Role'`,
  `'ClusterRole'`,
  `'RoleBinding'`,
  `'ClusterRoleBinding'`,
  `'CustomResourceDefinition'`,
  `'PodDisruptionBudget'`,
  `'HorizontalPodAutoscaler'`,
  `'PodSecurityPolicy'`,
  `'EndpointSlices'`,
]
export const defaultKeywords = [
  'select',
  'from',
  'where',
  'values',
  'as',
  'join',
  'on',
  'group by',
  'order by',
  'having',
  'or',
  'and',
]

export const tabsList = [
  { label: 'KeywordSearch', value: 'keyword', disabled: true },
  { label: 'SQLSearch', value: 'sql' },
  {
    label: 'SearchByNaturalLanguage',
    value: 'natural',
  },
]

export interface InsightTab {
  value: string
  label: string
  disabled?: boolean
}

export const insightTabsList: InsightTab[] = [
  { value: 'Topology', label: 'ResourceTopology' },
  { value: 'YAML', label: 'YAML.TabName' },
  { value: 'Events', label: 'EventAggregator' },
]

export const defaultSqlExamples = [
  `kind='Namespace'`,
  `kind!='Pod'`,
  `namespace='default'`,
  `cluster='democluster' and kind='Pod'`,
  `kind not in ('pod','service')`,
  `${"`annotations.app.kubernetes.io/name` = 'demoapp'"}`,
]

export const DEFALUT_PAGE_SIZE_10 = 10

export const Languages = [
  {
    label: 'English',
    value: 'en',
  },
  {
    label: '中文',
    value: 'zh',
  },
  {
    label: 'Deutsch',
    value: 'de',
  },
  {
    label: 'Português',
    value: 'pt',
  },
  {
    label: '한국어',
    value: 'ko',
  },
]

export const LanguagesMap = {
  zh: '简体中文',
  en: 'English',
  de: 'Deutsch',
  pt: 'Português',
  ko: '한국어',
}
