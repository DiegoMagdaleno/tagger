//
//  NSURLTagger.m
//  tagger
//
//  Created by Diego Magdaleno on 10/30/20.
//

#import <Foundation/Foundation.h>
#import "NSTaggerURL.h"


@implementation NSURL (Tags)
- (NSArray *) GetTags
{
    NSArray *currentTags = nil;
    NSError *tagError = nil;
    
    BOOL success = [self getResourceValue:&currentTags forKey:NSURLTagNamesKey error:&tagError];
    
    
    return currentTags;
}
@end
