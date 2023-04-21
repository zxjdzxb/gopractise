// Map 是一种无序的键值对的集合。
package main

import "fmt"

func main() {
	var siteMap map[string]string /*创建集合 */
	siteMap = make(map[string]string)

	/* map 插入 key - value 对,各个国家对应的首都 */
	siteMap["Google"] = "谷歌"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Baidu"] = "百度"
	siteMap["Wiki"] = "维基百科"
	fmt.Println(siteMap)
	/*使用键输出地图值 */
	for site := range siteMap {
		fmt.Println(site, "首都是", siteMap[site])
	}

	/*查看元素在集合中是否存在 */
	// name, ok := siteMap["Facebook"] /*如果确定是真实的,则存在,否则不存在 */
	name, ok := siteMap["Baidu"]
	/*fmt.Println(capital) */
	fmt.Println(name, ok)
	if ok {
		fmt.Println("Facebook 的 站点是", name)
	} else {
		fmt.Println("Facebook 站点不存在")
	}
	// 获取键值对
	v1 := siteMap["Facebook"]
	v2, ok := siteMap["Baidu"] // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值
	fmt.Println(v1, v2, ok)    // 输出： 0  false
	fmt.Println("======delete()======")
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*删除元素*/
	delete(countryCapitalMap, "France")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

}
